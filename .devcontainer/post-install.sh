#!/bin/bash -e

# setup git
# =====================================
git config --global --add safe.directory /workspaces

# setup twine
# ===========
pip install twine

# install pulumi
# ==============
curl -fsSL https://get.pulumi.com | sh

cat <<EOT >>"${HOME}/.profile"
if [ -d "\${HOME}/.pulumi/bin" ]; then
  PATH="\${HOME}/.pulumi/bin:\${PATH}"
fi
EOT

echo "source ${HOME}/.profile" >>"${HOME}/.bashrc"
echo "source ${HOME}/.profile" >>"${HOME}/.zshrc"


# install pulumictl
# =================
PULUMICTL_VERSION="0.0.32"

PULUMITCTL_URL="https://github.com/pulumi/pulumictl/releases/download/v${PULUMICTL_VERSION}/pulumictl-v${PULUMICTL_VERSION}-linux-$(dpkg --print-architecture).tar.gz"

wget -q "${PULUMITCTL_URL}" -O /tmp/pulumictl.tar.gz

tar -xzf /tmp/pulumictl.tar.gz -C /tmp

sudo mv /tmp/pulumictl /usr/local/bin/pulumictl
