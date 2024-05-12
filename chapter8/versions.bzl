"""Defines the dependencies' versions.

This allow us to keep the book generic and update cheaply the versions.
"""

# Go
# The last version can be found here: https://go.dev/dl/
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
GO_VERSION = "1.22.2"

# Rules_go
# The last version and SHA256 can be found here: https://github.com/bazelbuild/rules_go/releases
RULES_GO_VERSION = "v0.47.1"
RULES_GO_SHA256 = "f74c98d6df55217a36859c74b460e774abc0410a47cc100d822be34d5f990f16"

# Gazelle
# The last version and SHA256 can be found here: https://github.com/bazelbuild/bazel-gazelle/blob/master/README.rst
GAZELLE_VERSION = "v0.36.0"
GAZELLE_INTEGRITY = "sha256-dd8ojEsxyB61D1Hi4U9HY8t1SNquEmgXJHBkY3/Z6mI="

# Protobuf
# The last version can be found here: https://github.com/protocolbuffers/protobuf/releases
# Note: You do not need to download the last version,
#       Just update this variable. Bazel will download it for you.
PROTO_VERSION = "v26.1"

# protoc-gen-validate
# The last version can be found here: https://github.com/bufbuild/protoc-gen-validate/releases
PROTOC_GEN_VALIDATE_VERSION = "v1.0.4"
