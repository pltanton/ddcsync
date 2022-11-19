{ lib, stdenv, buildGoModule, makeWrapper, ddcutil }:
buildGoModule {
  pname = "ddcsync";
  version = "0";

  vendorSha256 =
    "6d297fdf7340ef4742cfd34c0cb7c29368cc804d84a564ffbc05578e4d10d0f0";

  src = lib.cleanSourceWith {
    filter = name: type:
      let baseName = baseNameOf (toString name);
      in !(lib.hasSuffix ".nix" baseName);
    src = lib.cleanSource ../.;
  };

  allowGoReference = true;

  nativeBuildInputs = [ makeWrapper ];

  postFixup = ''
    wrapProgram $out/bin/ddcsync --prefix PATH : ${lib.makeBinPath [ ddcutil ]}
  '';

  meta = with lib; {
    homepage = "https://github.com/pltanton/ddcsync";
    description = "Daemon to sync laptop brightness with DDS";
    license = licenses.bsd3;
    platforms = platforms.linux;
    mainProgram = "ddcutil";
  };
}
