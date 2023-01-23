#!/usr/bin/bash

#octave [-HVWdfhiqvx] [--debug] [--debug-jit] [--doc-cache-file file]
#       [--echo-commands] [--eval CODE] [--exec-path path]
#       [--gui] [--help] [--image-path path]
#       [--info-file file] [--info-program prog] [--interactive]
#       [--jit-compiler] [--line-editing] [--no-gui] [--no-history]
#       [--no-init-file] [--no-init-path] [--no-line-editing]
#       [--no-site-file] [--no-window-system] [--norc] [-p path]
#       [--path path] [--persist] [--silent] [--traditional]
#       [--verbose] [--version] [file]
#      see: https://docs.octave.org/v7.3.0/Command-Line-Options.html


fn_mfile='flt.m'
fn_m_log='m.log'

if [[ $# -eq 1 ]] ; then
        fn_mfile=$1
fi

tac $HOME/.octave_hist | awk '/^#/ {found++} ; found<2 ' |tac >$fn_m_log

octave $fn_mfile
