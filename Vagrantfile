VAGRANTFILE_API_VERSION = '2'

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # trusty (latest), virtualbox only
  config.vm.box = 'ubuntu/trusty64'

  config.vm.provider(:virtualbox) do |v|
    v.name = 'bosh-docker-cpi'
    v.customize ['modifyvm', :id, '--cpus', '4']
    v.customize ['modifyvm', :id, '--memory', '4096']
  end

  # mount bosh dir for testing
  config.vm.synced_folder('../', '/opt/bosh', owner: 'root', group: 'root')
  config.vm.provision('docker')

  config.vm.provision "shell", inline: 'echo "export PATH=$HOME/bin:$PATH" >> /home/vagrant/.bash_profile'
end