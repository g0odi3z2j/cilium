package ebpf

import (
	"bufio"
	"bytes"
	"debug/elf"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/cilium/ebpf/asm"
	"github.com/cilium/ebpf/internal"
	"github.com/cilium/ebpf/internal/btf"
	"github.com/cilium/ebpf/internal/unix"
)

// elfCode is a convenience to reduce the amount of arguments that have to
// be passed around explicitly. You should treat its contents as immutable.
type elfCode struct {
	*internal.SafeELFFile
	sections map[elf.SectionIndex]*elfSection
	license  string
	version  uint32
	btf      *btf.Spec
}

// LoadCollectionSpec parses an ELF file into a CollectionSpec.
func LoadCollectionSpec(file string) (*CollectionSpec, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	spec, err := LoadCollectionSpecFromReader(f)
	if err != nil {
		return nil, fmt.Errorf("file %s: %w", file, err)
	}
	return spec, nil
}

// LoadCollectionSpecFromReader parses an ELF file into a CollectionSpec.
func LoadCollectionSpecFromReader(rd io.ReaderAt) (*CollectionSpec, error) {
	f, err := internal.NewSafeELFFile(rd)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		licenseSection *elf.Section
		versionSection *elf.Section
		sections       = make(map[elf.SectionIndex]*elfSection)
		relSections    = make(map[elf.SectionIndex]*elf.Section)
	)

	// This is the target of relocations generated by inline assembly.
	sections[elf.SHN_UNDEF] = newElfSection(new(elf.Section), undefSection)

	// Collect all the sections we're interested in. This includes relocations
	// which we parse later.
	for i, sec := range f.Sections {
		idx := elf.SectionIndex(i)

		switch {
		case strings.HasPrefix(sec.Name, "license"):
			licenseSection = sec
		case strings.HasPrefix(sec.Name, "version"):
			versionSection = sec
		case strings.HasPrefix(sec.Name, "maps"):
			sections[idx] = newElfSection(sec, mapSection)
		case sec.Name == ".maps":
			sections[idx] = newElfSection(sec, btfMapSection)
		case sec.Name == ".bss" || sec.Name == ".data" || strings.HasPrefix(sec.Name, ".rodata"):
			sections[idx] = newElfSection(sec, dataSection)
		case sec.Type == elf.SHT_REL:
			// Store relocations under the section index of the target
			relSections[elf.SectionIndex(sec.Info)] = sec
		case sec.Type == elf.SHT_PROGBITS && (sec.Flags&elf.SHF_EXECINSTR) != 0 && sec.Size > 0:
			sections[idx] = newElfSection(sec, programSection)
		}
	}

	license, err := loadLicense(licenseSection)
	if err != nil {
		return nil, fmt.Errorf("load license: %w", err)
	}

	version, err := loadVersion(versionSection, f.ByteOrder)
	if err != nil {
		return nil, fmt.Errorf("load version: %w", err)
	}

	btfSpec, err := btf.LoadSpecFromReader(rd)
	if err != nil && !errors.Is(err, btf.ErrNotFound) {
		return nil, fmt.Errorf("load BTF: %w", err)
	}

	ec := &elfCode{
		SafeELFFile: f,
		sections:    sections,
		license:     license,
		version:     version,
		btf:         btfSpec,
	}

	symbols, err := f.Symbols()
	if err != nil {
		return nil, fmt.Errorf("load symbols: %v", err)
	}

	ec.assignSymbols(symbols)

	// Go through relocation sections, and parse the ones for sections we're
	// interested in. Make sure that relocations point at valid sections.
	for idx, relSection := range relSections {
		section := sections[idx]
		if section == nil {
			continue
		}

		rels, err := ec.loadRelocations(relSection, symbols)
		if err != nil {
			return nil, fmt.Errorf("relocation for section %q: %w", section.Name, err)
		}

		for _, rel := range rels {
			target := sections[rel.Section]
			if target == nil {
				return nil, fmt.Errorf("section %q: reference to %q in section %s: %w", section.Name, rel.Name, rel.Section, ErrNotSupported)
			}

			if target.Flags&elf.SHF_STRINGS > 0 {
				return nil, fmt.Errorf("section %q: string is not stack allocated: %w", section.Name, ErrNotSupported)
			}

			target.references++
		}

		section.relocations = rels
	}

	// Collect all the various ways to define maps.
	maps := make(map[string]*MapSpec)
	if err := ec.loadMaps(maps); err != nil {
		return nil, fmt.Errorf("load maps: %w", err)
	}

	if err := ec.loadBTFMaps(maps); err != nil {
		return nil, fmt.Errorf("load BTF maps: %w", err)
	}

	if err := ec.loadDataSections(maps); err != nil {
		return nil, fmt.Errorf("load data sections: %w", err)
	}

	// Finally, collect programs and link them.
	progs, err := ec.loadProgramSections()
	if err != nil {
		return nil, fmt.Errorf("load programs: %w", err)
	}

	return &CollectionSpec{maps, progs, ec.ByteOrder}, nil
}

func loadLicense(sec *elf.Section) (string, error) {
	if sec == nil {
		return "", nil
	}

	data, err := sec.Data()
	if err != nil {
		return "", fmt.Errorf("section %s: %v", sec.Name, err)
	}
	return string(bytes.TrimRight(data, "\000")), nil
}

func loadVersion(sec *elf.Section, bo binary.ByteOrder) (uint32, error) {
	if sec == nil {
		return 0, nil
	}

	var version uint32
	if err := binary.Read(sec.Open(), bo, &version); err != nil {
		return 0, fmt.Errorf("section %s: %v", sec.Name, err)
	}
	return version, nil
}

type elfSectionKind int

const (
	undefSection elfSectionKind = iota
	mapSection
	btfMapSection
	programSection
	dataSection
)

type elfSection struct {
	*elf.Section
	kind elfSectionKind
	// Offset from the start of the section to a symbol
	symbols map[uint64]elf.Symbol
	// Offset from the start of the section to a relocation, which points at
	// a symbol in another section.
	relocations map[uint64]elf.Symbol
	// The number of relocations pointing at this section.
	references int
}

func newElfSection(section *elf.Section, kind elfSectionKind) *elfSection {
	return &elfSection{
		section,
		kind,
		make(map[uint64]elf.Symbol),
		make(map[uint64]elf.Symbol),
		0,
	}
}

// assignSymbols takes a list of symbols and assigns them to their
// respective sections, indexed by name.
func (ec *elfCode) assignSymbols(symbols []elf.Symbol) {
	for _, symbol := range symbols {
		symType := elf.ST_TYPE(symbol.Info)
		symSection := ec.sections[symbol.Section]
		if symSection == nil {
			continue
		}

		// Anonymous symbols only occur in debug sections which we don't process
		// relocations for. Anonymous symbols are not referenced from other sections.
		if symbol.Name == "" {
			continue
		}

		// Older versions of LLVM don't tag symbols correctly, so keep
		// all NOTYPE ones.
		switch symSection.kind {
		case mapSection, btfMapSection, dataSection:
			if symType != elf.STT_NOTYPE && symType != elf.STT_OBJECT {
				continue
			}
		case programSection:
			if symType != elf.STT_NOTYPE && symType != elf.STT_FUNC {
				continue
			}
			// LLVM emits LBB_ (Local Basic Block) symbols that seem to be jump
			// targets within sections, but BPF has no use for them.
			if symType == elf.STT_NOTYPE && elf.ST_BIND(symbol.Info) == elf.STB_LOCAL &&
				strings.HasPrefix(symbol.Name, "LBB") {
				continue
			}
		// Only collect symbols that occur in program/maps/data sections.
		default:
			continue
		}

		symSection.symbols[symbol.Value] = symbol
	}
}

// loadProgramSections iterates ec's sections and emits a ProgramSpec
// for each function it finds.
//
// The resulting map is indexed by function name.
func (ec *elfCode) loadProgramSections() (map[string]*ProgramSpec, error) {

	progs := make(map[string]*ProgramSpec)

	// Generate a ProgramSpec for each function found in each program section.
	for _, sec := range ec.sections {
		if sec.kind != programSection {
			continue
		}

		if len(sec.symbols) == 0 {
			return nil, fmt.Errorf("section %v: missing symbols", sec.Name)
		}

		funcs, err := ec.loadFunctions(sec)
		if err != nil {
			return nil, fmt.Errorf("section %v: %w", sec.Name, err)
		}

		progType, attachType, progFlags, attachTo := getProgType(sec.Name)

		for name, insns := range funcs {
			spec := &ProgramSpec{
				Name:          name,
				Type:          progType,
				Flags:         progFlags,
				AttachType:    attachType,
				AttachTo:      attachTo,
				SectionName:   sec.Name,
				License:       ec.license,
				KernelVersion: ec.version,
				Instructions:  insns,
				ByteOrder:     ec.ByteOrder,
			}

			if ec.btf != nil {
				spec.BTF, err = ec.btf.Program(name)
				if err != nil && !errors.Is(err, btf.ErrNoExtendedInfo) {
					return nil, fmt.Errorf("program %s: %w", name, err)
				}
			}

			// Function names must be unique within a single ELF blob.
			if progs[name] != nil {
				return nil, fmt.Errorf("duplicate program name %s", name)
			}
			progs[name] = spec
		}
	}

	// Populate each prog's references with pointers to all of its callees.
	if err := populateReferences(progs); err != nil {
		return nil, fmt.Errorf("populating references: %w", err)
	}

	// Hide programs (e.g. library functions) that were not explicitly emitted
	// to an ELF section. These could be exposed in a separate CollectionSpec
	// field later to allow them to be modified.
	for n, p := range progs {
		if p.SectionName == ".text" {
			delete(progs, n)
		}
	}

	return progs, nil
}

// loadFunctions extracts instruction streams from the given program section
// starting at each symbol in the section. The section's symbols must already
// be narrowed down to STT_NOTYPE (emitted by clang <8) or STT_FUNC.
//
// The resulting map is indexed by function name.
func (ec *elfCode) loadFunctions(section *elfSection) (map[string]asm.Instructions, error) {
	var (
		r      = bufio.NewReader(section.Open())
		funcs  = make(map[string]asm.Instructions)
		offset uint64
		insns  asm.Instructions
	)
	for {
		ins := asm.Instruction{
			// Symbols denote the first instruction of a function body.
			Symbol: section.symbols[offset].Name,
		}

		// Pull one instruction from the instruction stream.
		n, err := ins.Unmarshal(r, ec.ByteOrder)
		if errors.Is(err, io.EOF) {
			fn := insns.Name()
			if fn == "" {
				return nil, errors.New("reached EOF before finding a valid symbol")
			}

			// Reached the end of the section and the decoded instruction buffer
			// contains at least one valid instruction belonging to a function.
			// Store the result and stop processing instructions.
			funcs[fn] = insns
			break
		}
		if err != nil {
			return nil, fmt.Errorf("offset %d: %w", offset, err)
		}

		// Decoded the first instruction of a function body but insns already
		// holds a valid instruction stream. Store the result and flush insns.
		if ins.Symbol != "" && insns.Name() != "" {
			funcs[insns.Name()] = insns
			insns = nil
		}

		if rel, ok := section.relocations[offset]; ok {
			// A relocation was found for the current offset. Apply it to the insn.
			if err = ec.relocateInstruction(&ins, rel); err != nil {
				return nil, fmt.Errorf("offset %d: relocate instruction: %w", offset, err)
			}
		} else {
			// Up to LLVM 9, calls to subprograms within the same ELF section are
			// sometimes encoded using relative jumps without relocation entries.
			// If, after all relocations entries have been processed, there are
			// still relative pseudocalls left, they must point to an existing
			// symbol within the section.
			// When splitting sections into subprograms, the targets of these calls
			// are no longer in scope, so they must be resolved here.
			if ins.IsFunctionReference() && ins.Constant != -1 {
				tgt := jumpTarget(offset, ins)
				sym := section.symbols[tgt].Name
				if sym == "" {
					return nil, fmt.Errorf("offset %d: no jump target found at offset %d", offset, tgt)
				}

				ins.Reference = sym
				ins.Constant = -1
			}
		}

		insns = append(insns, ins)
		offset += n
	}

	return funcs, nil
}

// jumpTarget takes ins' offset within an instruction stream (in bytes)
// and returns its absolute jump destination (in bytes) within the
// instruction stream.
func jumpTarget(offset uint64, ins asm.Instruction) uint64 {
	// A relative jump instruction describes the amount of raw BPF instructions
	// to jump, convert the offset into bytes.
	dest := ins.Constant * asm.InstructionSize

	// The starting point of the jump is the end of the current instruction.
	dest += int64(offset + asm.InstructionSize)

	if dest < 0 {
		return 0
	}

	return uint64(dest)
}

func (ec *elfCode) relocateInstruction(ins *asm.Instruction, rel elf.Symbol) error {
	var (
		typ  = elf.ST_TYPE(rel.Info)
		bind = elf.ST_BIND(rel.Info)
		name = rel.Name
	)

	target := ec.sections[rel.Section]

	switch target.kind {
	case mapSection, btfMapSection:
		if bind != elf.STB_GLOBAL {
			return fmt.Errorf("possible erroneous static qualifier on map definition: found reference to %q", name)
		}

		if typ != elf.STT_OBJECT && typ != elf.STT_NOTYPE {
			// STT_NOTYPE is generated on clang < 8 which doesn't tag
			// relocations appropriately.
			return fmt.Errorf("map load: incorrect relocation type %v", typ)
		}

		ins.Src = asm.PseudoMapFD

		// Mark the instruction as needing an update when creating the
		// collection.
		if err := ins.RewriteMapPtr(-1); err != nil {
			return err
		}

	case dataSection:
		var offset uint32
		switch typ {
		case elf.STT_SECTION:
			if bind != elf.STB_LOCAL {
				return fmt.Errorf("direct load: %s: unsupported relocation %s", name, bind)
			}

			// This is really a reference to a static symbol, which clang doesn't
			// emit a symbol table entry for. Instead it encodes the offset in
			// the instruction itself.
			offset = uint32(uint64(ins.Constant))

		case elf.STT_OBJECT:
			if bind != elf.STB_GLOBAL {
				return fmt.Errorf("direct load: %s: unsupported relocation %s", name, bind)
			}

			offset = uint32(rel.Value)

		default:
			return fmt.Errorf("incorrect relocation type %v for direct map load", typ)
		}

		// We rely on using the name of the data section as the reference. It
		// would be nicer to keep the real name in case of an STT_OBJECT, but
		// it's not clear how to encode that into Instruction.
		name = target.Name

		// The kernel expects the offset in the second basic BPF instruction.
		ins.Constant = int64(uint64(offset) << 32)
		ins.Src = asm.PseudoMapValue

		// Mark the instruction as needing an update when creating the
		// collection.
		if err := ins.RewriteMapPtr(-1); err != nil {
			return err
		}

	case programSection:
		switch opCode := ins.OpCode; {
		case opCode.JumpOp() == asm.Call:
			if ins.Src != asm.PseudoCall {
				return fmt.Errorf("call: %s: incorrect source register", name)
			}

			switch typ {
			case elf.STT_NOTYPE, elf.STT_FUNC:
				if bind != elf.STB_GLOBAL {
					return fmt.Errorf("call: %s: unsupported binding: %s", name, bind)
				}

			case elf.STT_SECTION:
				if bind != elf.STB_LOCAL {
					return fmt.Errorf("call: %s: unsupported binding: %s", name, bind)
				}

				// The function we want to call is in the indicated section,
				// at the offset encoded in the instruction itself. Reverse
				// the calculation to find the real function we're looking for.
				// A value of -1 references the first instruction in the section.
				offset := int64(int32(ins.Constant)+1) * asm.InstructionSize
				sym, ok := target.symbols[uint64(offset)]
				if !ok {
					return fmt.Errorf("call: no symbol at offset %d", offset)
				}

				name = sym.Name
				ins.Constant = -1

			default:
				return fmt.Errorf("call: %s: invalid symbol type %s", name, typ)
			}
		case opCode.IsDWordLoad():
			switch typ {
			case elf.STT_FUNC:
				if bind != elf.STB_GLOBAL {
					return fmt.Errorf("load: %s: unsupported binding: %s", name, bind)
				}

			case elf.STT_SECTION:
				if bind != elf.STB_LOCAL {
					return fmt.Errorf("load: %s: unsupported binding: %s", name, bind)
				}

				// ins.Constant already contains the offset in bytes from the
				// start of the section. This is different than a call to a
				// static function.

			default:
				return fmt.Errorf("load: %s: invalid symbol type %s", name, typ)
			}

			sym, ok := target.symbols[uint64(ins.Constant)]
			if !ok {
				return fmt.Errorf("load: no symbol at offset %d", ins.Constant)
			}

			name = sym.Name
			ins.Constant = -1
			ins.Src = asm.PseudoFunc

		default:
			return fmt.Errorf("neither a call nor a load instruction: %v", ins)
		}

	case undefSection:
		if bind != elf.STB_GLOBAL {
			return fmt.Errorf("asm relocation: %s: unsupported binding: %s", name, bind)
		}

		if typ != elf.STT_NOTYPE {
			return fmt.Errorf("asm relocation: %s: unsupported type %s", name, typ)
		}

		// There is nothing to do here but set ins.Reference.

	default:
		return fmt.Errorf("relocation to %q: %w", target.Name, ErrNotSupported)
	}

	ins.Reference = name
	return nil
}

func (ec *elfCode) loadMaps(maps map[string]*MapSpec) error {
	for _, sec := range ec.sections {
		if sec.kind != mapSection {
			continue
		}

		nSym := len(sec.symbols)
		if nSym == 0 {
			return fmt.Errorf("section %v: no symbols", sec.Name)
		}

		if sec.Size%uint64(nSym) != 0 {
			return fmt.Errorf("section %v: map descriptors are not of equal size", sec.Name)
		}

		var (
			r    = bufio.NewReader(sec.Open())
			size = sec.Size / uint64(nSym)
		)
		for i, offset := 0, uint64(0); i < nSym; i, offset = i+1, offset+size {
			mapSym, ok := sec.symbols[offset]
			if !ok {
				return fmt.Errorf("section %s: missing symbol for map at offset %d", sec.Name, offset)
			}

			mapName := mapSym.Name
			if maps[mapName] != nil {
				return fmt.Errorf("section %v: map %v already exists", sec.Name, mapSym)
			}

			lr := io.LimitReader(r, int64(size))

			spec := MapSpec{
				Name: SanitizeName(mapName, -1),
			}
			switch {
			case binary.Read(lr, ec.ByteOrder, &spec.Type) != nil:
				return fmt.Errorf("map %s: missing type", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.KeySize) != nil:
				return fmt.Errorf("map %s: missing key size", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.ValueSize) != nil:
				return fmt.Errorf("map %s: missing value size", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.MaxEntries) != nil:
				return fmt.Errorf("map %s: missing max entries", mapName)
			case binary.Read(lr, ec.ByteOrder, &spec.Flags) != nil:
				return fmt.Errorf("map %s: missing flags", mapName)
			}

			extra, err := io.ReadAll(lr)
			if err != nil {
				return fmt.Errorf("map %s: reading map tail: %w", mapName, err)
			}
			if len(extra) > 0 {
				spec.Extra = bytes.NewReader(extra)
			}

			if err := spec.clampPerfEventArraySize(); err != nil {
				return fmt.Errorf("map %s: %w", mapName, err)
			}

			maps[mapName] = &spec
		}
	}

	return nil
}

// loadBTFMaps iterates over all ELF sections marked as BTF map sections
// (like .maps) and parses them into MapSpecs. Dump the .maps section and
// any relocations with `readelf -x .maps -r <elf_file>`.
func (ec *elfCode) loadBTFMaps(maps map[string]*MapSpec) error {
	for _, sec := range ec.sections {
		if sec.kind != btfMapSection {
			continue
		}

		if ec.btf == nil {
			return fmt.Errorf("missing BTF")
		}

		// Each section must appear as a DataSec in the ELF's BTF blob.
		var ds *btf.Datasec
		if err := ec.btf.TypeByName(sec.Name, &ds); err != nil {
			return fmt.Errorf("cannot find section '%s' in BTF: %w", sec.Name, err)
		}

		// Open a Reader to the ELF's raw section bytes so we can assert that all
		// of them are zero on a per-map (per-Var) basis. For now, the section's
		// sole purpose is to receive relocations, so all must be zero.
		rs := sec.Open()

		for _, vs := range ds.Vars {
			// BPF maps are declared as and assigned to global variables,
			// so iterate over each Var in the DataSec and validate their types.
			v, ok := vs.Type.(*btf.Var)
			if !ok {
				return fmt.Errorf("section %v: unexpected type %s", sec.Name, vs.Type)
			}
			name := string(v.Name)

			// The BTF metadata for each Var contains the full length of the map
			// declaration, so read the corresponding amount of bytes from the ELF.
			// This way, we can pinpoint which map declaration contains unexpected
			// (and therefore unsupported) data.
			_, err := io.Copy(internal.DiscardZeroes{}, io.LimitReader(rs, int64(vs.Size)))
			if err != nil {
				return fmt.Errorf("section %v: map %s: initializing BTF map definitions: %w", sec.Name, name, internal.ErrNotSupported)
			}

			if maps[name] != nil {
				return fmt.Errorf("section %v: map %s already exists", sec.Name, name)
			}

			// Each Var representing a BTF map definition contains a Struct.
			mapStruct, ok := v.Type.(*btf.Struct)
			if !ok {
				return fmt.Errorf("expected struct, got %s", v.Type)
			}

			mapSpec, err := mapSpecFromBTF(sec, &vs, mapStruct, ec.btf, name, false)
			if err != nil {
				return fmt.Errorf("map %v: %w", name, err)
			}

			if err := mapSpec.clampPerfEventArraySize(); err != nil {
				return fmt.Errorf("map %v: %w", name, err)
			}

			maps[name] = mapSpec
		}

		// Drain the ELF section reader to make sure all bytes are accounted for
		// with BTF metadata.
		i, err := io.Copy(io.Discard, rs)
		if err != nil {
			return fmt.Errorf("section %v: unexpected error reading remainder of ELF section: %w", sec.Name, err)
		}
		if i > 0 {
			return fmt.Errorf("section %v: %d unexpected remaining bytes in ELF section, invalid BTF?", sec.Name, i)
		}
	}

	return nil
}

// mapSpecFromBTF produces a MapSpec based on a btf.Struct def representing
// a BTF map definition. The name and spec arguments will be copied to the
// resulting MapSpec, and inner must be true on any resursive invocations.
func mapSpecFromBTF(es *elfSection, vs *btf.VarSecinfo, def *btf.Struct, spec *btf.Spec, name string, inner bool) (*MapSpec, error) {
	var (
		key, value         btf.Type
		keySize, valueSize uint32
		mapType            MapType
		flags, maxEntries  uint32
		pinType            PinType
		innerMapSpec       *MapSpec
		contents           []MapKV
		err                error
	)

	for i, member := range def.Members {
		switch member.Name {
		case "type":
			mt, err := uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get type: %w", err)
			}
			mapType = MapType(mt)

		case "map_flags":
			flags, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF map flags: %w", err)
			}

		case "max_entries":
			maxEntries, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF map max entries: %w", err)
			}

		case "key":
			if keySize != 0 {
				return nil, errors.New("both key and key_size given")
			}

			pk, ok := member.Type.(*btf.Pointer)
			if !ok {
				return nil, fmt.Errorf("key type is not a pointer: %T", member.Type)
			}

			key = pk.Target

			size, err := btf.Sizeof(pk.Target)
			if err != nil {
				return nil, fmt.Errorf("can't get size of BTF key: %w", err)
			}

			keySize = uint32(size)

		case "value":
			if valueSize != 0 {
				return nil, errors.New("both value and value_size given")
			}

			vk, ok := member.Type.(*btf.Pointer)
			if !ok {
				return nil, fmt.Errorf("value type is not a pointer: %T", member.Type)
			}

			value = vk.Target

			size, err := btf.Sizeof(vk.Target)
			if err != nil {
				return nil, fmt.Errorf("can't get size of BTF value: %w", err)
			}

			valueSize = uint32(size)

		case "key_size":
			// Key needs to be nil and keySize needs to be 0 for key_size to be
			// considered a valid member.
			if key != nil || keySize != 0 {
				return nil, errors.New("both key and key_size given")
			}

			keySize, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF key size: %w", err)
			}

		case "value_size":
			// Value needs to be nil and valueSize needs to be 0 for value_size to be
			// considered a valid member.
			if value != nil || valueSize != 0 {
				return nil, errors.New("both value and value_size given")
			}

			valueSize, err = uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get BTF value size: %w", err)
			}

		case "pinning":
			if inner {
				return nil, errors.New("inner maps can't be pinned")
			}

			pinning, err := uintFromBTF(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't get pinning: %w", err)
			}

			pinType = PinType(pinning)

		case "values":
			// The 'values' field in BTF map definitions is used for declaring map
			// value types that are references to other BPF objects, like other maps
			// or programs. It is always expected to be an array of pointers.
			if i != len(def.Members)-1 {
				return nil, errors.New("'values' must be the last member in a BTF map definition")
			}

			if valueSize != 0 && valueSize != 4 {
				return nil, errors.New("value_size must be 0 or 4")
			}
			valueSize = 4

			valueType, err := resolveBTFArrayMacro(member.Type)
			if err != nil {
				return nil, fmt.Errorf("can't resolve type of member 'values': %w", err)
			}

			switch t := valueType.(type) {
			case *btf.Struct:
				// The values member pointing to an array of structs means we're expecting
				// a map-in-map declaration.
				if mapType != ArrayOfMaps && mapType != HashOfMaps {
					return nil, errors.New("outer map needs to be an array or a hash of maps")
				}
				if inner {
					return nil, fmt.Errorf("nested inner maps are not supported")
				}

				// This inner map spec is used as a map template, but it needs to be
				// created as a traditional map before it can be used to do so.
				// libbpf names the inner map template '<outer_name>.inner', but we
				// opted for _inner to simplify validation logic. (dots only supported
				// on kernels 5.2 and up)
				// Pass the BTF spec from the parent object, since both parent and
				// child must be created from the same BTF blob (on kernels that support BTF).
				innerMapSpec, err = mapSpecFromBTF(es, vs, t, spec, name+"_inner", true)
				if err != nil {
					return nil, fmt.Errorf("can't parse BTF map definition of inner map: %w", err)
				}

			case *btf.FuncProto:
				// The values member contains an array of function pointers, meaning an
				// autopopulated PROG_ARRAY.
				if mapType != ProgramArray {
					return nil, errors.New("map needs to be a program array")
				}

			default:
				return nil, fmt.Errorf("unsupported value type %q in 'values' field", t)
			}

			contents, err = resolveBTFValuesContents(es, vs, member)
			if err != nil {
				return nil, fmt.Errorf("resolving values contents: %w", err)
			}

		default:
			return nil, fmt.Errorf("unrecognized field %s in BTF map definition", member.Name)
		}
	}

	if key == nil {
		key = &btf.Void{}
	}
	if value == nil {
		value = &btf.Void{}
	}

	return &MapSpec{
		Name:       SanitizeName(name, -1),
		Type:       MapType(mapType),
		KeySize:    keySize,
		ValueSize:  valueSize,
		MaxEntries: maxEntries,
		Flags:      flags,
		BTF:        &btf.Map{Spec: spec, Key: key, Value: value},
		Pinning:    pinType,
		InnerMap:   innerMapSpec,
		Contents:   contents,
	}, nil
}

// uintFromBTF resolves the __uint macro, which is a pointer to a sized
// array, e.g. for int (*foo)[10], this function will return 10.
func uintFromBTF(typ btf.Type) (uint32, error) {
	ptr, ok := typ.(*btf.Pointer)
	if !ok {
		return 0, fmt.Errorf("not a pointer: %v", typ)
	}

	arr, ok := ptr.Target.(*btf.Array)
	if !ok {
		return 0, fmt.Errorf("not a pointer to array: %v", typ)
	}

	return arr.Nelems, nil
}

// resolveBTFArrayMacro resolves the __array macro, which declares an array
// of pointers to a given type. This function returns the target Type of
// the pointers in the array.
func resolveBTFArrayMacro(typ btf.Type) (btf.Type, error) {
	arr, ok := typ.(*btf.Array)
	if !ok {
		return nil, fmt.Errorf("not an array: %v", typ)
	}

	ptr, ok := arr.Type.(*btf.Pointer)
	if !ok {
		return nil, fmt.Errorf("not an array of pointers: %v", typ)
	}

	return ptr.Target, nil
}

// resolveBTFValuesContents resolves relocations into ELF sections belonging
// to btf.VarSecinfo's. This can be used on the 'values' member in BTF map
// definitions to extract static declarations of map contents.
func resolveBTFValuesContents(es *elfSection, vs *btf.VarSecinfo, member btf.Member) ([]MapKV, error) {
	// The elements of a .values pointer array are not encoded in BTF.
	// Instead, relocations are generated into each array index.
	// However, it's possible to leave certain array indices empty, so all
	// indices' offsets need to be checked for emitted relocations.

	// The offset of the 'values' member within the _struct_ (in bits)
	// is the starting point of the array. Convert to bytes. Add VarSecinfo
	// offset to get the absolute position in the ELF blob.
	start := (member.OffsetBits / 8) + vs.Offset
	// 'values' is encoded in BTF as a zero (variable) length struct
	// member, and its contents run until the end of the VarSecinfo.
	// Add VarSecinfo offset to get the absolute position in the ELF blob.
	end := vs.Size + vs.Offset
	// The size of an address in this section. This determines the width of
	// an index in the array.
	align := uint32(es.SectionHeader.Addralign)

	// Check if variable-length section is aligned.
	if (end-start)%align != 0 {
		return nil, errors.New("unaligned static values section")
	}
	elems := (end - start) / align

	if elems == 0 {
		return nil, nil
	}

	contents := make([]MapKV, 0, elems)

	// k is the array index, off is its corresponding ELF section offset.
	for k, off := uint32(0), start; k < elems; k, off = k+1, off+align {
		r, ok := es.relocations[uint64(off)]
		if !ok {
			continue
		}

		// Relocation exists for the current offset in the ELF section.
		// Emit a value stub based on the type of relocation to be replaced by
		// a real fd later in the pipeline before populating the map.
		// Map keys are encoded in MapKV entries, so empty array indices are
		// skipped here.
		switch t := elf.ST_TYPE(r.Info); t {
		case elf.STT_FUNC:
			contents = append(contents, MapKV{uint32(k), r.Name})
		case elf.STT_OBJECT:
			contents = append(contents, MapKV{uint32(k), r.Name})
		default:
			return nil, fmt.Errorf("unknown relocation type %v", t)
		}
	}

	return contents, nil
}

func (ec *elfCode) loadDataSections(maps map[string]*MapSpec) error {
	for _, sec := range ec.sections {
		if sec.kind != dataSection {
			continue
		}

		if sec.references == 0 {
			// Prune data sections which are not referenced by any
			// instructions.
			continue
		}

		if ec.btf == nil {
			return errors.New("data sections require BTF, make sure all consts are marked as static")
		}

		var datasec *btf.Datasec
		if err := ec.btf.TypeByName(sec.Name, &datasec); err != nil {
			return fmt.Errorf("data section %s: can't get BTF: %w", sec.Name, err)
		}

		data, err := sec.Data()
		if err != nil {
			return fmt.Errorf("data section %s: can't get contents: %w", sec.Name, err)
		}

		if uint64(len(data)) > math.MaxUint32 {
			return fmt.Errorf("data section %s: contents exceed maximum size", sec.Name)
		}

		mapSpec := &MapSpec{
			Name:       SanitizeName(sec.Name, -1),
			Type:       Array,
			KeySize:    4,
			ValueSize:  uint32(len(data)),
			MaxEntries: 1,
			Contents:   []MapKV{{uint32(0), data}},
			BTF:        &btf.Map{Spec: ec.btf, Key: &btf.Void{}, Value: datasec},
		}

		switch sec.Name {
		case ".rodata":
			mapSpec.Flags = unix.BPF_F_RDONLY_PROG
			mapSpec.Freeze = true
		case ".bss":
			// The kernel already zero-initializes the map
			mapSpec.Contents = nil
		}

		maps[sec.Name] = mapSpec
	}
	return nil
}

func getProgType(sectionName string) (ProgramType, AttachType, uint32, string) {
	types := []struct {
		prefix     string
		progType   ProgramType
		attachType AttachType
		progFlags  uint32
	}{
		// Please update the types from libbpf.c and follow the order of it.
		// https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/tools/lib/bpf/libbpf.c
		{"socket", SocketFilter, AttachNone, 0},
		{"sk_reuseport/migrate", SkReuseport, AttachSkReuseportSelectOrMigrate, 0},
		{"sk_reuseport", SkReuseport, AttachSkReuseportSelect, 0},
		{"kprobe/", Kprobe, AttachNone, 0},
		{"uprobe/", Kprobe, AttachNone, 0},
		{"kretprobe/", Kprobe, AttachNone, 0},
		{"uretprobe/", Kprobe, AttachNone, 0},
		{"tc", SchedCLS, AttachNone, 0},
		{"classifier", SchedCLS, AttachNone, 0},
		{"action", SchedACT, AttachNone, 0},
		{"tracepoint/", TracePoint, AttachNone, 0},
		{"tp/", TracePoint, AttachNone, 0},
		{"raw_tracepoint/", RawTracepoint, AttachNone, 0},
		{"raw_tp/", RawTracepoint, AttachNone, 0},
		{"raw_tracepoint.w/", RawTracepointWritable, AttachNone, 0},
		{"raw_tp.w/", RawTracepointWritable, AttachNone, 0},
		{"tp_btf/", Tracing, AttachTraceRawTp, 0},
		{"fentry/", Tracing, AttachTraceFEntry, 0},
		{"fmod_ret/", Tracing, AttachModifyReturn, 0},
		{"fexit/", Tracing, AttachTraceFExit, 0},
		{"fentry.s/", Tracing, AttachTraceFEntry, unix.BPF_F_SLEEPABLE},
		{"fmod_ret.s/", Tracing, AttachModifyReturn, unix.BPF_F_SLEEPABLE},
		{"fexit.s/", Tracing, AttachTraceFExit, unix.BPF_F_SLEEPABLE},
		{"freplace/", Extension, AttachNone, 0},
		{"lsm/", LSM, AttachLSMMac, 0},
		{"lsm.s/", LSM, AttachLSMMac, unix.BPF_F_SLEEPABLE},
		{"iter/", Tracing, AttachTraceIter, 0},
		{"syscall", Syscall, AttachNone, 0},
		{"xdp_devmap/", XDP, AttachXDPDevMap, 0},
		{"xdp_cpumap/", XDP, AttachXDPCPUMap, 0},
		{"xdp", XDP, AttachNone, 0},
		{"perf_event", PerfEvent, AttachNone, 0},
		{"lwt_in", LWTIn, AttachNone, 0},
		{"lwt_out", LWTOut, AttachNone, 0},
		{"lwt_xmit", LWTXmit, AttachNone, 0},
		{"lwt_seg6local", LWTSeg6Local, AttachNone, 0},
		{"cgroup_skb/ingress", CGroupSKB, AttachCGroupInetIngress, 0},
		{"cgroup_skb/egress", CGroupSKB, AttachCGroupInetEgress, 0},
		{"cgroup/skb", CGroupSKB, AttachNone, 0},
		{"cgroup/sock_create", CGroupSKB, AttachCGroupInetSockCreate, 0},
		{"cgroup/sock_release", CGroupSKB, AttachCgroupInetSockRelease, 0},
		{"cgroup/sock", CGroupSock, AttachCGroupInetSockCreate, 0},
		{"cgroup/post_bind4", CGroupSock, AttachCGroupInet4PostBind, 0},
		{"cgroup/post_bind6", CGroupSock, AttachCGroupInet6PostBind, 0},
		{"cgroup/dev", CGroupDevice, AttachCGroupDevice, 0},
		{"sockops", SockOps, AttachCGroupSockOps, 0},
		{"sk_skb/stream_parser", SkSKB, AttachSkSKBStreamParser, 0},
		{"sk_skb/stream_verdict", SkSKB, AttachSkSKBStreamVerdict, 0},
		{"sk_skb", SkSKB, AttachNone, 0},
		{"sk_msg", SkMsg, AttachSkMsgVerdict, 0},
		{"lirc_mode2", LircMode2, AttachLircMode2, 0},
		{"flow_dissector", FlowDissector, AttachFlowDissector, 0},
		{"cgroup/bind4", CGroupSockAddr, AttachCGroupInet4Bind, 0},
		{"cgroup/bind6", CGroupSockAddr, AttachCGroupInet6Bind, 0},
		{"cgroup/connect4", CGroupSockAddr, AttachCGroupInet4Connect, 0},
		{"cgroup/connect6", CGroupSockAddr, AttachCGroupInet6Connect, 0},
		{"cgroup/sendmsg4", CGroupSockAddr, AttachCGroupUDP4Sendmsg, 0},
		{"cgroup/sendmsg6", CGroupSockAddr, AttachCGroupUDP6Sendmsg, 0},
		{"cgroup/recvmsg4", CGroupSockAddr, AttachCGroupUDP4Recvmsg, 0},
		{"cgroup/recvmsg6", CGroupSockAddr, AttachCGroupUDP6Recvmsg, 0},
		{"cgroup/getpeername4", CGroupSockAddr, AttachCgroupInet4GetPeername, 0},
		{"cgroup/getpeername6", CGroupSockAddr, AttachCgroupInet6GetPeername, 0},
		{"cgroup/getsockname4", CGroupSockAddr, AttachCgroupInet4GetSockname, 0},
		{"cgroup/getsockname6", CGroupSockAddr, AttachCgroupInet6GetSockname, 0},
		{"cgroup/sysctl", CGroupSysctl, AttachCGroupSysctl, 0},
		{"cgroup/getsockopt", CGroupSockopt, AttachCGroupGetsockopt, 0},
		{"cgroup/setsockopt", CGroupSockopt, AttachCGroupSetsockopt, 0},
		{"struct_ops+", StructOps, AttachNone, 0},
		{"sk_lookup/", SkLookup, AttachSkLookup, 0},

		{"seccomp", SocketFilter, AttachNone, 0},
	}

	for _, t := range types {
		if !strings.HasPrefix(sectionName, t.prefix) {
			continue
		}

		if !strings.HasSuffix(t.prefix, "/") {
			return t.progType, t.attachType, t.progFlags, ""
		}

		return t.progType, t.attachType, t.progFlags, sectionName[len(t.prefix):]
	}

	return UnspecifiedProgram, AttachNone, 0, ""
}

func (ec *elfCode) loadRelocations(sec *elf.Section, symbols []elf.Symbol) (map[uint64]elf.Symbol, error) {
	rels := make(map[uint64]elf.Symbol)

	if sec.Entsize < 16 {
		return nil, fmt.Errorf("section %s: relocations are less than 16 bytes", sec.Name)
	}

	r := bufio.NewReader(sec.Open())
	for off := uint64(0); off < sec.Size; off += sec.Entsize {
		ent := io.LimitReader(r, int64(sec.Entsize))

		var rel elf.Rel64
		if binary.Read(ent, ec.ByteOrder, &rel) != nil {
			return nil, fmt.Errorf("can't parse relocation at offset %v", off)
		}

		symNo := int(elf.R_SYM64(rel.Info) - 1)
		if symNo >= len(symbols) {
			return nil, fmt.Errorf("offset %d: symbol %d doesn't exist", off, symNo)
		}

		symbol := symbols[symNo]
		rels[rel.Off] = symbol
	}

	return rels, nil
}
