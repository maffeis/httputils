pipeline {
    agent any
    tools {
        go 'go-1.13.4'
    }
    environment {
        GO111MODULE = 'on'
        GOPROXY = 'direct'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            environment {
                CODECOV_TOKEN = credentials('codecov_token')
            }
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
                sh "curl -s https://codecov.io/bash | bash -s -"
            }
        }
        stage('Code Lint') {
            steps {
                // To install golangci-lint: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b . v1.21.0
                // If you are getting a bogus error about parallel execution of golangci-lint, check for the existence 
                // of an obsolete /tmp/golangci-lint.lock
                
                sh '/mnt/jenkins/tools/go/golangci-lint run'
            }
        }
        stage('Release') {
            when {
                buildingTag()
            }
            environment {
                GITHUB_TOKEN = credentials('github_token')
            }
            steps {
                sh 'curl -sL https://git.io/goreleaser | bash'
            }
        }
    }
}