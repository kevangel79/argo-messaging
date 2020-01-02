pipeline {
    agent {
        docker {
            image 'epel7go'
            args '-u root'
            label 'slave02'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh './build-script.sh'
            }
        }
    }
}