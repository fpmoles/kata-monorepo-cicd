def projectList = []
projectList.add(new ProjectDescriptor("alpha", ProjectType.GOLANG, BuildType.DOCKER))
projectList.add(new ProjectDescriptor("beta", ProjectType.GOLANG, BuildType.DOCKER))
projectList.add(new ProjectDescriptor("gamma", ProjectType.GOLANG, BuildType.DOCKER))

pipeline{
    agent any
    stages{
        stage ('Test'){
            steps{
                script{
                    for(int i=0;i<projectList.size();i++){
                        switch(projectList[i].projectType){
                            case GOLANG:
                                println "GOLANG: " + projectList[i].name
                                break
                            case JAVASCRIPT:
                                println "JAVASCRIPT: " + projectList[i].name
                                break
                            case JAVA:
                                println "JAVA: " + projectList[i].name
                                break
                            case STATIC:
                                println "STATIC: " + projectList[i].name
                                break
                        }
                    }
                }
            }
        }
        stage('Build'){
            steps{
                script{
                    for(int i=0;i<projectList.size();i++){
                        switch(projectList[i].buildType){
                            case SHELL:
                                println "SHELL: " + projectList[i].name
                                break
                            case DOCKER:
                                println "DOCKER: " + projectList[i].name
                                break
                        }
                    }
                }
            }
        }
    }
}


class ProjectDescriptor{
    String name
    ProjectType projectType
    BuildType buildType
    boolean passedTests
    boolean built

    ProjectDescriptor(name, projectType, buildType){
        this.name = name
        this.projectType = projectType
        this.buildType = buildType
    }

}

enum ProjectType{
    GOLANG, JAVASCRIPT, JAVA, STATIC
}

enum BuildType{
    SHELL, DOCKER
}
