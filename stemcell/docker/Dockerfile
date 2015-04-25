FROM bosh:stemcell

COPY assets/agent.json /var/vcap/bosh/agent.json
COPY assets/metadata.json /var/vcap/bosh/metatada.json
COPY assets/userdata.json /var/vcap/bosh/userdata.json
COPY assets/settings.json /var/vcap/bosh/settings_source.json

COPY assets/bosh-agent /var/vcap/bosh/bin/bosh-agent
COPY assets/run_agent.sh /etc/service/agent/run
RUN chmod +x /etc/service/agent/run