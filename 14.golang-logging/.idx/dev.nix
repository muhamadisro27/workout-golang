# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  # Which nixpkgs channel to use.
  channel = "stable-23.11"; # or "unstable"
  # Use https://search.nixos.org/packages to find packages
  packages = [
    pkgs.go
    pkgs.mysql80
    pkgs.openssl
    pkgs.wire
  ];
  # Sets environment variables in the workspace
  env = { };
  idx = {
    # Search for the extensions you want on https://open-vsx.org/ and use "publisher.id"
    extensions = [
      "golang.go"
    ];
    # Enable previews and customize configuration
    previews = {
      enable = true;
      previews = {
        web = {
          command = [
            "nodemon"
            "--signal"
            "SIGHUP"
            "-w"
            "."
            "-e"
            "go,html"
            "-x"
            "go run server.go -addr localhost:$PORT"
          ];
          manager = "web";
        };
      };
    };
  };
}
