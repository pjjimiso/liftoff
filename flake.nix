{
  description = "Liftoff - a simple launch countdown timer and CLI written in Go";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?reg=nixos-unstable";
  };

  outputs = { self, nixpkgs, ... } @ inputs:
    let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in
    {

      packages.x86_64-linux.liftoff = pkgs.buildGoModule {
        pname = "liftoff";
        version = self.shortRev or "dirty";
        src = ./.;
        vendorHash = "sha256-n+UuXOybCdy/IWNoDuF7dPv/1mjmeFoje7qPXRnmPaM=";
      };

      devShells.x86_64-linux.default = pkgs.mkShell {
        buildInputs = with pkgs; [ go gopls gotools go-tools ];
      };

      packages.x86_64-linux.default = self.packages.x86_64-linux.liftoff;

      meta = {
        description = "A simple launch countdown timer and CLI";
        homepage = "https://github.com/pjjimiso/liftoff";
        mainProgram = "liftoff.go";
      };
    };
}
