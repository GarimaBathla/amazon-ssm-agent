#!/usr/bin/env bash

cp ${BGO_SPACE}/Tools/src/update/linux/install.sh ${BGO_SPACE}/bin/linux_amd64/
cp ${BGO_SPACE}/Tools/src/update/linux/uninstall.sh ${BGO_SPACE}/bin/linux_amd64/

chmod 755 ${BGO_SPACE}/bin/linux_amd64/install.sh ${BGO_SPACE}/bin/linux_amd64/uninstall.sh

tar -zcvf ${BGO_SPACE}/bin/updates/amazon-ssm-agent/`cat ${BGO_SPACE}/VERSION`/amazon-ssm-agent-linux-amd64.tar.gz  -C ${BGO_SPACE}/bin/linux_amd64/ amazon-ssm-agent.rpm install.sh uninstall.sh

tar -zcvf ${BGO_SPACE}/bin/updates/amazon-ssm-agent-updater/`cat ${BGO_SPACE}/VERSION`/amazon-ssm-agent-updater-linux-amd64.tar.gz  -C ${BGO_SPACE}/bin/linux_amd64/ updater

rm ${BGO_SPACE}/bin/linux_amd64/install.sh
rm ${BGO_SPACE}/bin/linux_amd64/uninstall.sh

