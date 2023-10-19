{ pkgs, ... }:

{
  packages = [
    pkgs.git
  ];

  languages.go.enable = true;
  languages.go.package = pkgs.go_1_21;

  scripts.run-tests.exec = ''
    go test -v $DEVENV_ROOT/...
  '';
}
