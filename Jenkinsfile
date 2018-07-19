#!/usr/bin/env groovy

def projects = []

try{

    node{

        stage('Checkout'){
            def scmVars = checkout scm
        }
        stage('Collect Details'){
            projects = getProjects()
        }
        stage('Test Artifacts'){

        }
        stage('Build Artifacts'){

        }
        stage('Deploy Artifacts'){

        }
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
    }
}
