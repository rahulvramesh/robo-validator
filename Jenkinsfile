#!groovy

pipeline {
    agent any

    environment {
        DISABLE_AUTH = 'true'
        DB_ENGINE    = 'sqlite'
    }

    stages {
        try {
            stage('Build') {
                steps {
                    sh 'printenv'
                }
            }
        }catch(error) {
            throw error
        } finally {
        }
    }
}
