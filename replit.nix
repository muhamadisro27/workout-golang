{ pkgs }: {
  deps = [
    pkgs.mysql
    pkgs.mysql80
    pkgs.go_1_20
    pkgs.cowsay
  ];
}