package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/noironetworks/cilium-net/common/types"

	"github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/gorilla/mux"
)

func (router *Router) ping(w http.ResponseWriter, r *http.Request) {
	if str, err := router.daemon.Ping(); err != nil {
		processServerError(w, r, err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, str)
	}
}

func (router *Router) endpointCreate(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var ep types.Endpoint
	if err := d.Decode(&ep); err != nil {
		processServerError(w, r, err)
		return
	}
	if err := router.daemon.EndpointJoin(ep); err != nil {
		processServerError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (router *Router) endpointDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//FIXME: change uuid to a better designation
	if val, ok := vars["uuid"]; !ok {
		processServerError(w, r, errors.New("server received empty uuid"))
		return
	} else {
		if err := router.daemon.EndpointLeave(val); err != nil {
			processServerError(w, r, err)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func (router *Router) allocateIPv6(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if containerID, ok := vars["containerID"]; !ok {
		processServerError(w, r, errors.New("server received empty containerID"))
		return
	} else {
		ipamConfig, err := router.daemon.AllocateIPs(containerID)
		if err != nil {
			processServerError(w, r, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		e := json.NewEncoder(w)
		if err := e.Encode(ipamConfig); err != nil {
			processServerError(w, r, err)
			return
		}
	}
}

func (router *Router) releaseIPv6(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if containerID, ok := vars["containerID"]; !ok {
		processServerError(w, r, errors.New("server received empty containerID"))
		return
	} else {
		if err := router.daemon.ReleaseIPs(containerID); err != nil {
			processServerError(w, r, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (router *Router) getLabels(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["uuid"]
	if !exists {
		processServerError(w, r, errors.New("server received empty labels UUID"))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		processServerError(w, r, fmt.Errorf("server received invalid UUID '%s': '%s'", idStr, err))
	}
	labels, err := router.daemon.GetLabels(id)
	if err != nil {
		processServerError(w, r, err)
		return
	}
	if labels == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	if err := e.Encode(labels); err != nil {
		processServerError(w, r, err)
		return
	}
}

func (router *Router) getLabelsID(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var labels types.Labels
	if err := d.Decode(&labels); err != nil {
		processServerError(w, r, err)
		return
	}
	id, err := router.daemon.GetLabelsID(labels)
	if err != nil {
		processServerError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	lr := types.LabelsResponse{
		ID: id,
	}
	e := json.NewEncoder(w)
	if err := e.Encode(lr); err != nil {
		processServerError(w, r, err)
		return
	}
}
