def projects = []
pipeline{
    agent any
    stages{
        stage('Prepare'){
            steps{
                script{
                    def rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
                    println rawResults
                    projects = getProjects(rawResults)
                }
            }
        }
        stage('Test'){
            tools{
                go 'go'
            }
            steps{
                script{
                    for(Project project in projects){
                        executeTests(project)
                    }
                }
            }
        }
        stage('Build'){
            steps{
                script{
                    for(Project project in projects){
                        buildProject(project, ${env.BRANCH_NAME}, ${env.BUILD_NUMBER})
                    }
                }
            }
        }
    }
}


class Project{
    String name
    boolean testStatus
    String tag


    Project(String name){
        this.name = name
    }

    void setTagValue(String branchName, String buildNumber){
        if(branchName.equals('master')){
            this.tag = "release." + buildNumber
        }else if(branchName.contains('integration')){
            this.tag = "integration." + buildNumber
        }else{
            String newName = branchName.split("_")[0]
            this.tag="feature." + newName + "." + buildNumber
        }
    }

    @Override
    public String toString(){
        String result = "[Project: {"
        result = result + "name=" + this.name
        result = result + ", testStatus=" + this.testStatus
        result = result + "}]"
        return result
    }

}

Project[] getProjects(String rawResults){
    def results = []
    def tempProjects = []
    tempProjects = rawResults.split("\n")
    for(String project in tempProjects){
        results.add(new Project(project.trim()))
    }
    return results
}

void executeTests(Project project){
    String resultString = sh(returnStdout: true, script: "cd ${project.name} && bin/test.sh")
    boolean result = resultString.toBoolean()
    project.testStatus = result
}

void buildProject(Project project, String branchName, String buildNumber){
    project.setTagValue(branchName, buildNumber)
    println project.toString()
}


