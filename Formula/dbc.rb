# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Dbc < Formula
  desc "Database Connect"
  homepage "https://github.com/birdicare/homebrew-dbc"
  version "0.3.17"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/birdiecare/homebrew-dbc/releases/download/v0.3.17/birdiecare_dbc_0.3.17_darwin_arm64.tar.gz"
      sha256 "784d251739a9a3be6c1c9a52acc052d0f0423ddad05699dda25995e206724aa7"

      def install
        bin.install "dbc"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/birdiecare/homebrew-dbc/releases/download/v0.3.17/birdiecare_dbc_0.3.17_x86_64_arm64.tar.gz"
      sha256 "601bcd3162a789e94b269f2822b8bf5c2f179c6ea5a6d86e2615707aee7dd1a8"

      def install
        bin.install "dbc"
      end
    end
  end

  def post_install
    Kernel.system "install_ssm_plugin.sh"
  end
end
