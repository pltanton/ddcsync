{
  description =
    "DDCSync is daemon, that syncs laptop monitor brightness with external monitor";

  inputs = { nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable"; };

  outputs = inputs@{ self, nixpkgs, ... }:
    let
      inherit (nixpkgs) lib;
      genSystems = lib.genAttrs [
        # Add more systems if they are supported
        "aarch64-linux"
        "x86_64-linux"
      ];

      pkgsFor = genSystems (system: import nixpkgs { inherit system; });

      mkDate = longDate:
        (lib.concatStringsSep "-" [
          (builtins.substring 0 4 longDate)
          (builtins.substring 4 2 longDate)
          (builtins.substring 6 2 longDate)
        ]);
    in {
      overlays.default = _: prev: rec {
        ddcsync = prev.callPackage ./nix/default.nix { };
      };

      packages = genSystems (system:
        (self.overlays.default null pkgsFor.${system}) // {
          default = self.packages.${system}.ddcsync;
        });

      devShells = genSystems (system: {
        default = pkgsFor.${system}.mkShell {
          name = "ddcsync-shell";
          inputsFrom = [ self.packages.${system}.ddcsync ];
        };
      });

      formatter = genSystems (system: pkgsFor.${system}.alejandra);

      # nixosModules.default = import ./nix/module.nix self;
      homeManagerModules.default = import ./nix/hm-module.nix self;
    };
}
