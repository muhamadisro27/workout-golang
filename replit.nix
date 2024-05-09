{ pkgs }: {
  deps = [
    pkgs.mysql-shell
    pkgs.redis
    pkgs.mysql
    pkgs.mysql80
    pkgs.go_1_20
    pkgs.cowsay
  ];
}