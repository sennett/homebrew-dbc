install_plugin()
{
    echo "Installing AWS SSM plugin"

    curl "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/mac_arm64/sessionmanager-bundle.zip" -o "sessionmanager-bundle.zip"

    unzip sessionmanager-bundle.zip

    sudo ./sessionmanager-bundle/install -i /usr/local/sessionmanagerplugin -b /usr/local/bin/session-manager-plugin

    rm -rf ./sessionmanager-bundle
    rm -rf sessionmanager-bundle.zip
}


[ ! -d ~/../../usr/local/sessionmanagerplugin ] && install_plugin