ubuntu_mirror = 'http://archive.ubuntu.com/ubuntu/'
# ubuntu_mirror = 'http://mirror.rcg.sfu.ca/mirror/ubuntu/'
ubuntu_release = 'focal'
ubuntu_version = '20.04'
username = 'vagrant'
user_home = '/home/' + username
project_home = user_home + '/project/mq-demos' # you may need to change the working directory to match your project


python3_packages = '/usr/local/lib/python3.8/dist-packages'
ruby_gems = '/var/lib/gems/2.7.0/gems/'


# Get Ubuntu sources set up and packages up to date.

template '/etc/apt/sources.list' do
  variables(
    :mirror => ubuntu_mirror,
    :release => ubuntu_release
  )
  notifies :run, 'execute[apt-get update]', :immediately
end
execute 'apt-get update' do
  action :nothing
end
execute 'apt-get upgrade' do
  command 'apt-get dist-upgrade -y'
  only_if 'apt list --upgradeable | grep -q upgradable'
end
directory '/opt'
directory '/opt/installers'


# Basic packages many of us probably want. Includes gcc C and C++ compilers.

#package ['build-essential', 'cmake']


# Other core language tools you might want

#package ['python3', 'python3-pip', 'python3-dev']  # Python
#package ['ghc', 'libghc-random-dev', 'cabal-install']  # Haskell
#package 'golang-go'  # Go
#package 'erlang'  # Erlang
#package 'ocaml-nox'  # OCaml
#package ['rustc', 'cargo']  # Rust
#package 'scala'  # Scala 2.11
#package ['ruby', 'ruby-dev']  # Ruby
#package ['openjdk-11-jdk', 'maven']  # Java
#package ['php-cli', 'php-pear']  # PHP
#package 'clang' # Clang C/C++ compiler


# Scala 2.13
# prerequisite: a Java runtime, openjdk-11-jdk or similar
#scala_version = '2.13.3'
#remote_file '/opt/installers/scala.deb' do
#  # download URL for *.deb from https://scala-lang.org/download/
#  source "https://downloads.lightbend.com/scala/#{scala_version}/scala-#{scala_version}.deb"
#end
#execute 'dpkg -i /opt/installers/scala.deb' do
#  creates '/usr/bin/scala'
#end

# SBT
#execute 'sbt apt key' do
#  command 'curl -sL "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x2EE0EA64E40A89B84B2DF73499E82A75642AC823" | apt-key add'
#  not_if 'apt-key list | grep "2EE0 EA64"'
#end
#file '/etc/apt/sources.list.d/sbt.list' do
#  content 'deb https://dl.bintray.com/sbt/debian /'
#  notifies :run, 'execute[apt-get update]', :immediately
#end
#package 'sbt'


# .NET Core

#remote_file '/opt/installers/packages-microsoft-prod.deb' do
#  source "https://packages.microsoft.com/config/ubuntu/#{ubuntu_version}/packages-microsoft-prod.deb"
#end
#execute 'dpkg -i /opt/installers/packages-microsoft-prod.deb' do
#  creates '/etc/apt/sources.list.d/microsoft-prod.list'
#  notifies :run, 'execute[apt-get update]', :immediately
#end
#package ['dotnet-sdk-3.1']


# NodeJS (more modern than Ubuntu nodejs package) and NPM

#remote_file '/opt/installers/node-setup.sh' do
#  source 'https://deb.nodesource.com/setup_14.x'
#  mode '0755'
#end
#execute '/opt/installers/node-setup.sh' do
#  creates '/etc/apt/sources.list.d/nodesource.list'
#  notifies :run, 'execute[apt-get update]', :immediately
#end
#package ['nodejs']


# Go (more modern than Ubuntu golang-go package)

#execute 'snap install --classic go' do
#end

# SWIG

#package 'swig'


# RabbitMQ-related things

#package ['rabbitmq-server']

# Python pika library
#execute 'pip3 install pika==1.1.0' do
#  creates "#{python3_packages}/pika/__init__.py"
#end
# Ruby bunny library
#execute 'gem install bunny -v 2.17.0' do
#  creates "#{ruby_gems}/bunny-2.17.0/Gemfile"
#end
# Go amqp library
#execute 'go get github.com/streadway/amqp github.com/google/uuid' do
#  cwd project_home 
#  user username
#  environment 'HOME' => user_home
#  creates user_home + '/go/src/github.com/streadway/amqp/README.md'
#end
# Java amqp library
#package 'librabbitmq-client-java'


# ZeroMQ-related things

# C/C++ library and dev library
#package ['libzmq5', 'libzmq5-dev']
# Python pyzmq library
#execute 'pip3 install pyzmq==19.0.1' do
#  creates "#{python3_packages}/zmq/__init__.py"
#end
# Ruby ezmq library
#execute 'gem install ezmq -v 0.4.12' do
#  creates "#{ruby_gems}/ezmq-0.4.12/Gemfile"
#end
# Node zmq library
#execute 'npm install zeromq@6.0.0-beta.6' do
#  cwd project_home
#  user username
#  environment 'HOME' => user_home
#  creates project_home + '/node_modules/zeromq/package.json'
#end
# Go zmq4 library
#package 'pkg-config'
#execute 'go get github.com/pebbe/zmq4' do
#  cwd project_home 
#  user username
#  environment 'HOME' => user_home
#  creates user_home + '/go/pkg/linux_amd64/github.com/pebbe/zmq4.a'
#end


# GraalVM

#graalvm_version = '20.2.0'
#graalvm_directory = "graalvm-ce-java11-#{graalvm_version}"
#remote_file '/opt/installers/graalvm.tar.gz' do
#  source "https://github.com/graalvm/graalvm-ce-builds/releases/download/vm-#{graalvm_version}/graalvm-ce-java11-linux-amd64-#{graalvm_version}.tar.gz"
#end
#execute 'tar zxf /opt/installers/graalvm.tar.gz' do
#  cwd '/opt'
#  creates "/opt/#{graalvm_directory}/release"
#end
#link '/opt/graalvm' do
#  to "/opt/#{graalvm_directory}"
#  link_type :symbolic
#end
#file '/etc/profile.d/graalvm.sh' do
#  content 'PATH=${PATH}:/opt/graalvm/bin'
#end
# GraalVM's Python
#execute '/opt/graalvm/bin/gu install python' do
#  creates "/opt/#{graalvm_directory}/bin/graalpython"
#end
# GraalVM's Ruby
#execute '/opt/graalvm/bin/gu install ruby' do
#  creates "/opt/#{graalvm_directory}/bin/ruby"
#end
# GraalVM's LLVM (C, C++) tools
#execute '/opt/graalvm/bin/gu install llvm-toolchain' do
#  creates "/opt/#{graalvm_directory}/bin/lli"
#end
