# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Dbc < Formula
  desc "Database Connect"
  homepage "https://github.com/birdicare/homebrew-dbc"
  version "0.5.1"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/birdiecare/homebrew-dbc/releases/download/v0.5.1/birdiecare_dbc_0.5.1_darwin_arm64.tar.gz"
      sha256 "75b8e3a23db2fea452ecd335d02b213990f5fc23592bf2cd493c896e7a1174a0"

      def install
        bin.install "dbc"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/birdiecare/homebrew-dbc/releases/download/v0.5.1/birdiecare_dbc_0.5.1_x86_64_arm64.tar.gz"
      sha256 "cbcf29e18ca3b0c2137d0981b43ce5f312e4ded4ed8eee31762ef79464278b66"

      def install
        bin.install "dbc"
      end
    end
  end
end
