self:
{ config, lib, pkgs, ... }:
with lib;
let cfg = config.services.ddcsync;
in {
  options.services.ddcsync = {
    enable = lib.mkEnableOption "daemon to sync external monitor brightness";
  };
  config = mkIf cfg.enable {
    systemd.user.services = {
      ddcsync = {
        Unit = {
          Description =
            "DDCSync: daemon to sync external monitor brightness service";
          PartOf = [ "graphical-session.target" ];
        };

        Service = {
          Type = "simple";
          ExecStart = "${pkgs.ddcsync}/bin/ddcsync";
          Restart = "always";
        };

        Install = { WantedBy = [ "graphical-session.target" ]; };
      };
    };
  };
}
