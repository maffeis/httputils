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
            steps {
                sh 'go test ./...'
            }
        }
    }
}