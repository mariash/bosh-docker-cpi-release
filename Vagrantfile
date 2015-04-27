VAGRANTFILE_API_VERSION = '2'

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = 'ubuntu/trusty64'

  config.vm.provider(:virtualbox) do |v|
    v.name = 'bosh-docker-cpi'
    v.customize ['modifyvm', :id, '--cpus', '4']
    v.customize ['modifyvm', :id, '--memory', '4096']
  end

  config.vm.provision('docker')
end
