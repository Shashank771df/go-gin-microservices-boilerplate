for dir in ./*
do
    fdir=${dir%*/}
    deploy=../deploys/${fdir##*/}
    if [ -e ${fdir}/start.go ]
    then
        if [ -d ${deploy} ]
        then
            go build ${fdir##*/}/*.go
            cp ${fdir##*/}/.env ${deploy}
            mv start ${deploy}
            sha=`sha1sum ${deploy}/start`
            echo "Created an application ${fdir##*/} with SHA1 ${sha}"
        else
            echo "Directory ${deploy} does not exists"
        fi
    else
        echo "Directorio ${fdir} no tiene archivo start.go"
    fi
done