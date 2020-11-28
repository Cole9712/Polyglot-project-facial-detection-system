ubuntu_mirror = 'http://archive.ubuntu.com/ubuntu/'
# ubuntu_mirror = 'http://mirror.rcg.sfu.ca/mirror/ubuntu/'
ubuntu_release = 'focal'
ubuntu_version = '20.04'
username = 'vagrant'
user_home = '/home/' + username
project_home = user_home + '/project/' # you may need to change the working directory to match your project


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

package ['build-essential', 'cmake']


# Other core language tools you might want

package ['python3', 'python3-pip', 'python3-dev']  # Python
#package ['ghc', 'libghc-random-dev', 'cabal-install']  # Haskell
# package 'golang-go'  # Go
#package 'erlang'  # Erlang
#package 'ocaml-nox'  # OCaml
#package ['rustc', 'cargo']  # Rust
#package 'scala'  # Scala 2.11
#package ['ruby', 'ruby-dev']  # Ruby
#package ['openjdk-11-jdk', 'maven']  # Java
#package ['php-cli', 'php-pear']  # PHP
#package 'clang' # Clang C/C++ compiler


# NodeJS (more modern than Ubuntu nodejs package) and NPM

remote_file '/opt/installers/node-setup.sh' do
 source 'https://deb.nodesource.com/setup_14.x'
 mode '0755'
end
execute '/opt/installers/node-setup.sh' do
 creates '/etc/apt/sources.list.d/nodesource.list'
 notifies :run, 'execute[apt-get update]', :immediately
end
package ['nodejs']


# Go (more modern than Ubuntu golang-go package)

execute 'snap install --classic go' do
end

directory '/project/client/'
execute 'npm install --save-dev node-sass' do
end
execute 'npm install holderjs' do
end
execute 'npm install fine-uploader' do
end
execute 'npm install' do
end
execute 'npm run serve' do
end



