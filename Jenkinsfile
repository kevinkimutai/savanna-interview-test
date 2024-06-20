pipeline {
    agent any
    
   tools { go 'go1.22.1' }
    
    stages {
        stage('Checkout') {
            steps {
                echo '--- Checking out the code from version control ---'
                // Checkout your code from version control (e.g., Git)
                 git branch: 'main', url: 'https://github.com/kevinkimutai/jenkins-project.git'
            }
        }
        
        stage('Build') {
            steps {
                echo '--- Building the GoLang application ---'
                // Build your GoLang application
                sh 'go build -o main ./cmd/main.go'
            }
        }
        
        stage('Test') {
            steps {
                echo '--- Running tests ---'
                // Run tests if any
                sh 'go test -cover ./...'
            }
        }

        stage('SonarQube analysis') {
            steps {
                script {
                    def scannerHome = tool 'sonarqube5.01'
                    withSonarQubeEnv('sonarserver') {
                        sh "${scannerHome}/bin/sonar-scanner \
                    -Dsonar.projectKey=savanna \
                    -Dsonar.sources=./ \
                    -Dsonar.go.coverage.reportPaths=coverage.out \
                    -Dsonar.go.tests.reportPaths=report.json"
                    }
                }
            }
        }


        
        stage('Deploy') {
            steps {
                echo '--- Deploying the application ---'
                // Deploy your application
                // You may use tools like Docker, Kubernetes, etc. for deployment
                sh 'echo "Deploying the application"'
                // Example: Deploy to Kubernetes
                // sh 'kubectl apply -f deployment.yaml'
            }
        }
    }
    
    post {
        always {
            echo '--- Cleaning up workspace ---'
            // Clean up workspace after build
            cleanWs()
        }
    }
}