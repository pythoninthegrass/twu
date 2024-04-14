# -*- mode: Python -*-

# ! naive
# docker_build('twu', '.', dockerfile='Dockerfile')
# k8s_yaml('kubernetes.yml')
# k8s_resource('twu', port_forwards=8080)

# * optimized
load('ext://restart_process', 'docker_build_with_restart')

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -installsuffix cgo -o build/twu .'

local_resource(
  'go-compile',
  compile_cmd,
  deps=['./main.go'])

docker_build_with_restart(
  'twu',
  '.',
  entrypoint=['/app/build/twu'],
  dockerfile='Dockerfile',
  only=[
    './build'
  ],
  live_update=[
    sync('./build', '/app/build')
  ],
)

k8s_yaml('kubernetes.yml')

k8s_resource('twu',
             port_forwards=8080,
             resource_deps=['go-compile']
)
