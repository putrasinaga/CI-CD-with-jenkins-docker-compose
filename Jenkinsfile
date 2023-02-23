pipeline{
    agent any

// parameters{
//     string(name:NAME, description:"silahkan masukan nama project")
// }
    stages{
        stage ("test"){
            steps{
                echo(
                    "testing"
                )
            }
        }
        stage("build project"){
            steps{
                echo("build image & run project")
                sh("docker compose build")
            }
        }
    }

    post{
        always{
            echo("proses sedang berjalan")
        }
        success{
            echo("proses berhasil")
        }
        failure{
            echo("proses gagal")
        }
       cleanup{
            echo("proses selesai")
        }
    }
}