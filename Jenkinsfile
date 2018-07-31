pipeline {
    agent any

    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }

    stages {
        stage('Build') {
            steps {
                sh 'printenv'
                
                if env.GIT_BRANCH == "origin/master" {
                    echo "master"
                }
            }
        }
    }
}
