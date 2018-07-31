#!groovy

pipeline {
    agent any
    
    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }

    stages {
            def app
            stage('Build') {
                steps {
                    sh 'printenv'
                }
            }
    }
}
