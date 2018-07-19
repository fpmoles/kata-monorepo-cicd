#!/usr/bin/env groovy

def projects = []

try{
    stage('Checkout'){
        node{
            def scmVars = checkout scm
            stash
        }
    }
    stage('Collect Details'){
        node{
            unstash
            projects = getProjects()
            stash
        }
    }
    stage('Test Artifacts'){
        node{
            unstash
            for(Project project in projects){
                result = testArtifact(project)
            }
            stash
        }
    }
    stage('Build Artifacts'){

    }
    stage('Deploy Artifacts'){

    }
}catch (exc){

}finally{

}


List getProjects(){
    sh "ls -l"
    rawResults=sh(returnStdout: true, script: "ls -l | egrep \'^d\' | awk \'{print \$9}\'")
    echo "Raw Results: ${rawResults}"
    List results=rawResults.split("\n")
    proj = []
    for(String result in results){
        Project project = new Project(result);
        println project.toString()
        proj.add(project)
    }
    return proj
}

boolean testProject(Project project){
    sh "cd ${project.name}"
    project.testsPass = sh (returnStdout: true, script: bin/test.sh)
    sh "cd .."
}

class Project{
    String name
    boolean needsBuild
    boolean testsPass

    Project(name){
        this.name = name
    }

    String toString(){
        String result = "[Project: {"
        result = result + "name=" + name;
        result = result + ", needsBuild=" + needsBuild
        result = result + ", testsPass=" + testsPass
        result = result + "}]"
        return result
    }
}
