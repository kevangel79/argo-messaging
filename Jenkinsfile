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
                sh 'echo Building >>>>>>>'
            }
        }
        stage('Test') {
            steps {
                sh './scripts/run-tests.sh'
            }
        }
    }
}