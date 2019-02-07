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
# import sys
import re
import subprocess
# sys.path.insert(0, os.path.abspath('.'))


# -- General configuration ------------------------------------------------

# If your documentation needs a minimal Sphinx version, state it here.
#
# needs_sphinx = '1.0'

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = ['sphinx.ext.ifconfig',
    'sphinx.ext.githubpages',
    'sphinx.ext.extlinks',
    'sphinxcontrib.openapi',
    'sphinx_tabs.tabs',
    'sphinxcontrib.spelling',
    'versionwarning.extension' ]

# Add any paths that contain templates here, relative to this directory.
templates_path = ['_templates']

# The suffix(es) of source filenames.
# You can specify multiple suffix as a list of string:
#
# source_suffix = ['.rst', '.md']
source_suffix = ['.rst', '.md']
source_parsers = {
   '.md': 'recommonmark.parser.CommonMarkParser',
}

# The master toctree document.
master_doc = 'index'

# General information about the project.
project = u'Cilium'
copyright = u'2017-2019, Cilium Authors'
author = u'Cilium Authors'

# The version info for the project you're documenting, acts as replacement for
# |version| and |release|, also used in various other places throughout the
# built documents.
#
# The short X.Y version.
release = open("../VERSION", "r").read().strip()
# Used by version warning
versionwarning_body_selector = "div.document"

# Fetch the docs version from an environment variable.
# Map latest -> master.
# Map stable -> current version number.
branch = os.environ.get('READTHEDOCS_VERSION')
if branch == None or branch == 'latest':
    branch = 'HEAD'
elif branch == 'stable':
    branch = release
githubusercontent = 'https://raw.githubusercontent.com/cilium/cilium/'
scm_web = githubusercontent + branch

# Store variables in the epilogue so they are globally available.
rst_epilog = """
.. |SCM_WEB| replace:: \{s}
.. |SCM_BRANCH| replace:: \{b}
""".format(s = scm_web, b = branch)

extlinks = {'git-tree': (scm_web + "/%s", '')}

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = None

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This patterns also effect to html_static_path and html_extra_path
exclude_patterns = ['_build', 'Thumbs.db', '.DS_Store', '_themes/**/*.rst']

# The name of the Pygments (syntax highlighting) style to use.
pygments_style = 'sphinx'

# If true, `todo` and `todoList` produce output, else they produce nothing.
todo_include_todos = False


# -- Options for HTML output ----------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme = "sphinx_rtd_theme"
html_theme_path = ["_themes/sphinx_rtd_theme", ]
html_style = "static/css/theme.css"
html_context = {
        'release': release
}

# Theme options are theme-specific and customize the look and feel of a theme
# further.  For a list of options available for each theme, see the
# documentation.
#
# html_theme_options = {}

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ['images', '_static', '_themes/sphinx_rtd_theme/sphinx_rtd_theme']

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
    app.add_stylesheet('parsed-literal.css')
    app.add_stylesheet('copybutton.css')
    app.add_javascript('clipboardjs.min.js')
    app.add_javascript("copybutton.js")
