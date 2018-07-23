def projects = []
pipeline{
    agent any
    stages{
        stage('Prepare'){
            steps{
                script{
                    projects = getProjects()
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
                        if(!project.testStatus){
                            currentBuild.result = 'UNSTABLE'
                        }
                    }
                }
            }
        }
        stage('Build'){
            //tools{
            //    docker 'docker'
            //}
            steps{
                script{
                    for(Project project in projects){
                        buildProject(project, "${env.BRANCH_NAME}", "${env.BUILD_NUMBER}")
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
    String buildVersion


    Project(String name){
        this.name = name
    }

    void setTagValue(String branchName, String buildNumber){
        println "In setTagValue"
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
        result = result + ", tag=" + this.tag
        result = result + ", buildVersion=" + this.buildVersion
        result = result + "}]"
        return result
    }

}

Project[] getProjects(){
    def results = []
    def tempProjects = []
    def tempResults = []
    def rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
    tempProjects = rawResults.split("\n")
    for(String project in tempProjects){
        println "Adding new project with name: " + project
        tempResults.add(new Project(project.trim()))
    }
    for(Project project in tempResults){
        println "Testing project with name: " + project.name
        String found = sh(returnStdOut: true, script: "if [ -f ${project.name}/Dockerfile ]; then echo true; else echo false; fi ")
        boolean exists = found.trim().toBoolean()
        if(exists){
            println "Adding project with name: " + project.name
            results.add(project)
        }
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
    String result = sh(returnStdout: true, script: "cd ${project.name} && bin/build.sh ${project.name} ${project.tag}")
    project.buildVersion = result.trim()
}


