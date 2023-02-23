pipeline{
    agent any

    tools{
        go '1.19'
    }

// parameters{
//     string(name:NAME, description:"silahkan masukan nama project")
// }
    stages{
        stage ("test"){
            steps{
                echo("testing")
                sh("go test -v")
                //sh("pwd")
            }
        }
        stage("build project"){
            steps{
                echo("build image")
                sh("docker compose build")
            }
        }

        stage("start container"){
            steps{
                echo("run container")
                sh("docker compose create && docker compose start")
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