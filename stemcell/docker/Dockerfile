FROM bosh:stemcell

COPY assets/userdata.json /var/vcap/bosh/warden-cpi-user-data.json
COPY assets/settings.json /var/vcap/bosh/settings_source.json

EXPOSE 6868
CMD /usr/sbin/runsvdir-start
