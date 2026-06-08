{
  description = "MusMGR's devShell flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    {
      nixpkgs,
      ...
    }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          biome
          curl
          docker-compose
          git
          go
          go-task
          goose
          jq
          lefthook
          nixpkgs-fmt
          nodejs_24
          pdfcpu
          pnpm
          postgresql_18
          python313Packages.sqlfmt
          sqlc
        ];

        shellHook = ''
          if [[ ! -f .git/hooks/pre-commit ]]; then
            echo "Installing lefthook hooks"
            lefthook install
          fi
        '';
      };
    };
}
