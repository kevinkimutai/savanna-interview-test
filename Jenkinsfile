pipeline {
    agent any

    tools { 
        go 'go1.22.1'
    }

    stages {
        stage('Checkout') {
            steps {
                echo '--- Checking out the code from version control ---'
                git branch: 'main', url: 'https://github.com/kevinkimutai/savanna-interview-test'
            }
        }
        
        stage('Build') {
            steps {
                echo '--- Building the GoLang application ---'
                withCredentials([file(credentialsId: 'env-file', variable: 'ENV_FILE')]) {
                    sh '''
                    set -o allexport
                    . $ENV_FILE
                    set -o allexport
                    go build -o main ./cmd/main.go
                    '''
                }
            }
        }
        
        stage('Test') {
            steps {
                echo '--- Running tests ---'
                withCredentials([file(credentialsId: 'env-file', variable: 'ENV_FILE')]) {
                    sh '''
                    set -o allexport
                    . $ENV_FILE
                    set -o allexport
                    go test -cover ./...
                    '''
                }
            }
        }

        // stage('SonarQube analysis') {
        //     steps {
        //         script {
        //            def scannerHome = tool 'sonarqube5.01'
        //             withSonarQubeEnv('sonarserver') {
        //                 withCredentials([file(credentialsId: 'env-file', variable: 'ENV_FILE')]) {
        //                     sh '''
        //                     set -o allexport
        //                     . $ENV_FILE
        //                     set -o allexport
        //                     . ${scannerHome}/bin/sonar-scanner \
        //                     -Dsonar.projectKey=savanna \
        //                     -Dsonar.sources=./ \
        //                     -Dsonar.go.coverage.reportPaths=coverage.out \
        //                     -Dsonar.go.tests.reportPaths=report.json
        //                     '''
        //                 }
        //             }
        //         }
        //     }
        // }

        environment {
        SCANNER_HOME = tool name: 'sonarqube5.01', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
        }
    
        stages {
            stage('SonarQube analysis') {
                steps {
                    script {
                        withSonarQubeEnv('sonarserver') {
                            withCredentials([file(credentialsId: 'env-file', variable: 'ENV_FILE')]) {
                                sh """
                                set -o allexport
                                . $ENV_FILE
                                ${env.SCANNER_HOME}/bin/sonar-scanner \
                                -Dsonar.projectKey=savanna \
                                -Dsonar.sources=./ \
                                -Dsonar.go.coverage.reportPaths=coverage.out \
                                -Dsonar.go.tests.reportPaths=report.json
                                """
                            }
                        }
                    }
                }
            }
        }
        
        stage('Deploy') {
            steps {
                echo '--- Deploying the application ---'
                withCredentials([file(credentialsId: 'env-file', variable: 'ENV_FILE')]) {
                    sh '''
                    set -o allexport
                    . $ENV_FILE
                    set -o allexport
                    echo "Deploying the application"
                    // Example: Deploy to Kubernetes
                    // sh 'kubectl apply -f deployment.yaml'
                    '''
                }
            }
        }
    }
    
    post {
        always {
            echo '--- Cleaning up workspace ---'
            cleanWs()
        }
    }
}
