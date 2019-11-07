pipeline {
    agent any
    tools {
        go 'go-1.13.4'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            environment {
                CODECOV_TOKEN = credentials('codecov_token_maffeis')
            }
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
                sh "curl -s https://codecov.io/bash | bash -s -"
            }
        }
        stage('Code Analysis') {
            steps {
                //sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b . v1.21.0'
                //sh 'echo hello world'
                //sh 'rm -f /tmp/golangci-lint.lock'
                sh '/mnt/jenkins/tools/go/golangci-lint run'
            }
        }
    }
}