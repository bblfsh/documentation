pipeline {
  agent {
    kubernetes {
      label 'sdk-drivers-updater'
      defaultContainer 'sdk-drivers-updater'
      yaml """
spec:
  nodeSelector:
    srcd.host/type: jenkins-worker
  affinity:
    podAntiAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchExpressions:
          - key: jenkins
            operator: In
            values:
            - slave
        topologyKey: kubernetes.io/hostname
  containers:
  - name: sdk-drivers-updater
    image: bblfsh/performance:latest
    imagePullPolicy: Always
    securityContext:
      privileged: true
    command:
    - dockerd
    tty: true
"""
    }
  }
  triggers {
    cron('0 0 * * 1')
    GenericTrigger(
      genericVariables: [
        [key: 'target', value: '$.target']
      ],
      token: 'update_languages',
      causeString: 'Triggered on $target',

      printContributedVariables: true,
      printPostContent: true,

      regexpFilterText: '$target',
      regexpFilterExpression: 'master'
    )
  }
  stages {
    stage('Run updater') {
      when { branch 'master' }
      steps {
        withCredentials([usernamePassword(credentialsId: '87b3cad8-8b12-4e91-8f47-33f3d7d45620', passwordVariable: 'token', usernameVariable: 'user')]) {
          sh 'GITHUB_TOKEN=${token} make update'
        }
      }
    }
  }
}
