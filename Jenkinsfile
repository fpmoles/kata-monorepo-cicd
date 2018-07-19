#!/usr/bin/env groovy

def String[] projects

try{
    stage('Collect Details'){
        projects = getProjects()
    }
    stage('Test Artifacts'){

    }
    stage('Build Artifacts'){

    }
    stage('Deploy Artifacts'){

    }
}catch (exc){

}finally{

}


Project[] getProjects(){
    node{
        rawResults=sh(returnStdout: true, script: "cd ${env.WORKSPACE} && ls -l | egrep \'^d\' | awk \'{print \$9}\'")
        echo "Raw Results: ${rawResults}"
        results[]=rawResults.split("\n").trim()
        for(int i=0;i<results.size();i++){

            result = results[i]
            echo "Results value for ${i}: ${result}"
        }
        return results;
    }
}

class Project{
    String name
    boolean needsBuild
    boolean testsPass

    Project(name){
        this.name = name
    }
}
