pipeline {
    agent {
        docker {
            image 'epel7go'
            args '-u root'
            label 'slave02'
        }
    }
    stages {
        stage ('Build') {
            steps {
                sh './scripts/build.sh'
            }
        }
        stage('Test') {
            steps {
                sh './scripts/run-tests.sh'
            }
        }
        stage ('Upload Artifact') {
            when{
                branch 'master'
            }
            steps {
                sh 'echo Upload Artifacts'
            }
        }
    }
}