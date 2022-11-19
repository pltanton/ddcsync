# Service to sync external monitor with laptop monitor by DDC

### Requirements

- ddcutil

## Build

Just build with go: `make build`, binary should be in `out/`

## Nixos installation

### With flakes

```nix
# flake.nix
{

  ddcsync.url = "path:/home/anton/Workdir/ddcsync";
  ddcsync.inputs.nixpkgs.follows = "nixpkgs";

  outputs = { self, nixpkgs, home-manager, ddcsync }: {
    modules = [
      ({ pkgs, ... }: {
        nixpkgs.overlays = [ ddcsync.overlays.default ]; # To use programm as package
      })
    ];

    # To use with home-manager
    homeConfigurations."USER@HOSTNAME"= home-manager.lib.homeManagerConfiguration {
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
      modules = [
        ddcsync.homeManagerModules.default
        { services.ddcsync.enable = true; }
        # ...
      ];
    };
  };

  # ...
```
