{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/057f9aecfb71c4437d2b27d3323df7f93c010b7e.tar.gz") {} }:

let
  go = pkgs.go_1_21;
in
pkgs.mkShell {
  buildInputs = [ go ];
}