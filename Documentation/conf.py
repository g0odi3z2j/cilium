# -*- coding: utf-8 -*-
#
# Cilium documentation build configuration file, created by
# sphinx-quickstart on Sun Feb 12 18:34:43 2017.
#
# This file is execfile()d with the current directory set to its
# containing dir.
#
# Note that not all possible configuration values are present in this
# autogenerated file.
#
# All configuration values have a default; values that are commented out
# serve to show the default.

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#
import os
import sys
import re
import subprocess
import semver

sys.path.insert(0, os.path.abspath('_exts'))
import cilium_external_links  # noqa: E402
import cilium_spellfilters  # noqa: E402

# -- General configuration ------------------------------------------------

# If your documentation needs a minimal Sphinx version, state it here.
#
# needs_sphinx = '1.0'

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
html_logo = "images/logo.svg"
extensions = ['myst_parser',
              'sphinx.ext.ifconfig',
              'sphinx.ext.githubpages',
              'sphinx.ext.extlinks',
              'sphinxcontrib.openapi',
              'sphinx_tabs.tabs',
              'sphinxcontrib.googleanalytics',
              'sphinxcontrib.spelling',
              'versionwarning.extension']

# Add any paths that contain templates here, relative to this directory.
templates_path = ['_templates']

# The suffix(es) of source filenames.
# You can specify multiple suffix as a list of string:
#
# source_suffix = ['.rst', '.md']
source_suffix = ['.rst', '.md']

# The master toctree document.
master_doc = 'index'

# General information about the project.
project = u'Cilium'
copyright = u'Cilium Authors'
author = u'Cilium Authors'

# The version info for the project you're documenting, acts as replacement for
# |version| and |release|, also used in various other places throughout the
# built documents.
#
# The short X.Y version.
release = open("../VERSION", "r").read().strip()
# Used by version warning
versionwarning_body_selector = "div.document"
versionwarning_api_url = "https://docs.cilium.io/"

# The version of Go used to compile Cilium
go_release = open("../GO_VERSION", "r").read().strip()

# The image tag for Cilium docker images
image_tag = 'v' + release

# Fetch the docs version from an environment variable.
# Map latest -> master.
# Map stable -> current version number.
branch = os.environ.get('READTHEDOCS_VERSION')
if not branch or branch == 'latest':
    branch = 'HEAD'
    archive_name = 'master'
    chart_release = './cilium'
    image_tag = 'latest'
elif branch == 'stable':
    branch = release
    archive_name = release
    chart_release = 'cilium/cilium --version ' + release
    tags.add('stable')
else:
    archive_name = branch
    chart_release = 'cilium/cilium --version ' + release
    tags.add('stable')
relinfo = semver.parse_version_info(release)
current_release = '%d.%d' % (relinfo.major, relinfo.minor)
if relinfo.patch == 90:
    next_release = '%d.%d' % (relinfo.major, relinfo.minor + 1)
    prev_release = current_release
else:
    next_release = current_release
    prev_release = '%d.%d' % (relinfo.major, relinfo.minor - 1)
githubusercontent = 'https://raw.githubusercontent.com/cilium/cilium/'
scm_web = githubusercontent + branch
jenkins_branch = 'https://jenkins.cilium.io/view/Cilium-v' + current_release
github_repo = 'https://github.com/cilium/cilium/'
archive_filename = archive_name + '.tar.gz'
archive_link = github_repo + 'archive/' + archive_filename
archive_name = 'cilium-' + archive_name.strip('v')
project_link = github_repo + 'projects?query=is:open+' + next_release
backport_format = github_repo + 'pulls?q=is:open+is:pr+-label:backport/author+label:%s/' + current_release

# Store variables in the epilogue so they are globally available.
rst_epilog = """
.. |SCM_WEB| replace:: \\{s}
.. |SCM_BRANCH| replace:: \\{b}
.. |SCM_ARCHIVE_NAME| replace:: \\{a}
.. |SCM_ARCHIVE_FILENAME| replace:: \\{f}
.. |SCM_ARCHIVE_LINK| replace:: \\{l}
.. |CURRENT_RELEASE| replace:: \\{c}
.. |NEXT_RELEASE| replace:: \\{n}
.. |CHART_RELEASE| replace:: \\{h}
.. |GO_RELEASE| replace:: \\{g}
.. |IMAGE_TAG| replace:: \\{i}
""".format(s=scm_web, b=branch, a=archive_name, f=archive_filename, l=archive_link, c=current_release, n=next_release, h=chart_release, g=go_release, i=image_tag)

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = "en"

extlinks = {
    'git-tree': (scm_web + "/%s", None),
    'jenkins-branch': (jenkins_branch + "/%s", None),
    'github-project': (project_link + '%s', None),
    'github-backport': (backport_format, None),
    'gh-issue': (github_repo + 'issues/%s', 'GitHub issue %s'),
    'prev-docs': (versionwarning_api_url + language + '/v' + prev_release + '/%s', None),
}

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This patterns also effect to html_static_path and html_extra_path
exclude_patterns = ['_build', 'Thumbs.db', '.DS_Store']

# The name of the Pygments (syntax highlighting) style to use.
pygments_style = 'sphinx'

# The default language to highlight source code in.
highlight_language = 'none'

# If true, `todo` and `todoList` produce output, else they produce nothing.
todo_include_todos = False

# Ignore spelling errors in generated files.
spelling_exclude_patterns = ['_api/v1/*/README.md']

# Add custom filters for spell checks.
spelling_filters = [cilium_spellfilters.WireGuardFilter]

# Ignore some warnings from MyST parser
suppress_warnings = ['myst.header']

googleanalytics_id = 'G-V9SYWYG92Y'

# -- Options for HTML output ----------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme = "sphinx_rtd_theme_cilium"

html_context = {
    'release': release
}

# Theme options are theme-specific and customize the look and feel of a theme
# further.  For a list of options available for each theme, see the
# documentation.
#
html_theme_options = {
    'logo_only': True
}

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ['images', '_static']

# Add any extra paths that contain custom files (such as robots.txt or
# .htaccess) here, relative to this directory. These files are copied
# directly to the root of the documentation.
html_extra_path = ['robots/robots.txt']

# -- Options for HTMLHelp output ------------------------------------------

# Output file base name for HTML help builder.
htmlhelp_basename = 'Ciliumdoc'


# -- Options for LaTeX output ---------------------------------------------

latex_elements = {
    # The paper size ('letterpaper' or 'a4paper').
    #
    # 'papersize': 'letterpaper',

    # The font size ('10pt', '11pt' or '12pt').
    #
    # 'pointsize': '10pt',

    # Additional stuff for the LaTeX preamble.
    #
    # 'preamble': '',

    # Latex figure (float) alignment
    #
    # 'figure_align': 'htbp',
    'extraclassoptions': 'openany',
}

# Grouping the document tree into LaTeX files. List of tuples
# (source start file, target name, title,
#  author, documentclass [howto, manual, or own class]).
latex_documents = [
    (master_doc, 'Cilium.tex', u'Cilium Documentation',
     u'Cilium Authors', 'manual'),
]


# -- Options for manual page output ---------------------------------------

# One entry per manual page. List of tuples
# (source start file, name, description, authors, manual section).
man_pages = [
    (master_doc, 'cilium', u'Cilium Documentation',
     [author], 1)
]


# -- Options for Texinfo output -------------------------------------------

# Grouping the document tree into Texinfo files. List of tuples
# (source start file, target name, title, author,
#  dir menu entry, description, category)
texinfo_documents = [
    (master_doc, 'Cilium', u'Cilium Documentation',
     author, 'Cilium', 'One line description of project.',
     'Miscellaneous'),
]

http_strict_mode = False

# Try as hard as possible to find references
default_role = 'any'


def setup(app):
    app.add_css_file('parsed-literal.css')
    app.add_css_file('copybutton.css')
    app.add_css_file('editbutton.css')
    app.add_js_file('clipboardjs.min.js')
    app.add_js_file("copybutton.js")
    app.add_css_file('helm-reference.css')
    # Patch HTML translator to open external links in new tabs
    app.set_translator("html", cilium_external_links.PatchedHTMLTranslator)
