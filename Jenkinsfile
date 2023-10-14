pipeline {
    agent any
    tools {
        go 'golang 1.21.3'
    }
    environment {
        CGO_ENABLED=1
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'make test'
            }
        }

        stage("code-coverage") {
            steps {
                echo 'CODE COVERAGE EXECUTION STARTED'
                sh 'make coverage'
            }
        }

        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go get ./...'
                sh 'docker build . -t jtelio/clean-architecture-go'
            }
        }

        stage('deliver') {
            agent any
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                    sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                    sh 'docker push jtelio/clean-architecture-go'
                }
            }
        }
    }
}
