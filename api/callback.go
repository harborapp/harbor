package api

import (
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// CallbackEvents collects registry events.
func CallbackEvents(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  http.StatusInternalServerError,
				"message": err,
			},
		)

		c.Abort()
		return
	}

	logrus.Debugf("%v", c.Params)
	logrus.Debugf("%v", c.Request.Header)
	logrus.Debugf("%v", string(b))

	c.JSON(
		http.StatusOK,
		"",
	)
}

// Headers: map[User-Agent:[Go-http-client/1.1] Content-Length:[1137] Content-Type:[application/vnd.docker.distribution.events.v1+json] Accept-Encoding:[gzip]]

// Push

// {
//    "events": [
//       {
//          "id": "d8e90701-133d-4e6a-b4b8-dd00dd4d78a9",
//          "timestamp": "2016-09-01T19:55:53.565651Z",
//          "action": "push",
//          "target": {
//             "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
//             "size": 505,
//             "digest": "sha256:a59906e33509d14c036c8678d687bd4eec81ed7c4b8ce907b888c607f6a1e0e6",
//             "length": 505,
//             "repository": "webhippie/busybox",
//             "url": "http://192.168.64.5:5000/v2/webhippie/busybox/manifests/sha256:a59906e33509d14c036c8678d687bd4eec81ed7c4b8ce907b888c607f6a1e0e6",
//             "tag": "latest"
//          },
//          "request": {
//             "id": "6fbb3ed1-8090-4dac-a796-ada940dedca0",
//             "addr": "192.168.64.5:45872",
//             "host": "192.168.64.5:5000",
//             "method": "PUT",
//             "useragent": "docker/1.10.0 go/go1.5.3 git-commit/590d5108 kernel/4.3.3-dhyve os/linux arch/amd64"
//          },
//          "actor": {},
//          "source": {
//             "addr": "712f606432fe:5000",
//             "instanceID": "9c70a6e9-f3c8-4994-ae9c-b74d045b66eb"
//          }
//       }
//    ]
// }

// Pull

// {
//    "events": [
//       {
//          "id": "405b552a-bba2-421c-ac76-6cf72a97197d",
//          "timestamp": "2016-09-01T19:55:53.555651Z",
//          "action": "pull",
//          "target": {
//             "mediaType": "application/octet-stream",
//             "size": 1459,
//             "digest": "sha256:2b8fd9751c4c0f5dd266fcae00707e67a2545ef34f9a29354585f93dac906749",
//             "length": 1459,
//             "repository": "webhippie/busybox",
//             "url": "http://192.168.64.5:5000/v2/webhippie/busybox/blobs/sha256:2b8fd9751c4c0f5dd266fcae00707e67a2545ef34f9a29354585f93dac906749"
//          },
//          "request": {
//             "id": "5784fe9c-774f-4329-8554-7a5fcb9b4c16",
//             "addr": "192.168.64.5:45868",
//             "host": "192.168.64.5:5000",
//             "method": "HEAD",
//             "useragent": "docker/1.10.0 go/go1.5.3 git-commit/590d5108 kernel/4.3.3-dhyve os/linux arch/amd64"
//          },
//          "actor": {},
//          "source": {
//             "addr": "712f606432fe:5000",
//             "instanceID": "9c70a6e9-f3c8-4994-ae9c-b74d045b66eb"
//          }
//       }
//    ]
// }

// Mount

// {
//    "events": [
//       {
//          "id": "8aab1d13-9f84-4f2f-a759-f7cdf40e54d8",
//          "timestamp": "2016-09-01T19:18:48.205121Z",
//          "action": "mount",
//          "target": {
//             "mediaType": "application/octet-stream",
//             "size": 667590,
//             "digest": "sha256:8ddc19f16526912237dd8af81971d5e4dd0587907234be2b83e249518d5b673f",
//             "length": 667590,
//             "repository": "busybox",
//             "fromRepository": "webhippie/busybox",
//             "url": "http://192.168.64.5:5000/v2/busybox/blobs/sha256:8ddc19f16526912237dd8af81971d5e4dd0587907234be2b83e249518d5b673f"
//          },
//          "request": {
//             "id": "f5e5d99c-7d43-4b0d-8844-a2e678ccebbc",
//             "addr": "192.168.64.5:45786",
//             "host": "192.168.64.5:5000",
//             "method": "POST",
//             "useragent": "docker/1.10.0 go/go1.5.3 git-commit/590d5108 kernel/4.3.3-dhyve os/linux arch/amd64"
//          },
//          "actor": {},
//          "source": {
//             "addr": "d9ddb2f6bae9:5000",
//             "instanceID": "c1ed9df0-9e3d-4080-b1dd-1ca6b7198090"
//          }
//       }
//    ]
// }

// Delete

// {
//    "events": [
//       {
//          "id": "60eef10b-33bf-4cf8-b967-3765e18a7199",
//          "timestamp": "2016-09-01T20:01:57.64765Z",
//          "action": "delete",
//          "target": {
//             "digest": "sha256:a59906e33509d14c036c8678d687bd4eec81ed7c4b8ce907b888c607f6a1e0e6",
//             "repository": "webhippie/busybox"
//          },
//          "request": {
//             "id": "43561bde-7799-4058-b581-47f7296d1243",
//             "addr": "192.168.64.1:59928",
//             "host": "192.168.64.5:5000",
//             "method": "DELETE",
//             "useragent": "HTTPie/0.9.6"
//          },
//          "actor": {},
//          "source": {
//             "addr": "712f606432fe:5000",
//             "instanceID": "9c70a6e9-f3c8-4994-ae9c-b74d045b66eb"
//          }
//       }
//    ]
// }

// # http GET http://192.168.64.5:5000/v2/_catalog
// HTTP/1.1 200 OK
// Content-Length: 39
// Content-Type: application/json; charset=utf-8
// Date: Thu, 01 Sep 2016 22:51:20 GMT
// Docker-Distribution-Api-Version: registry/2.0
// X-Content-Type-Options: nosniff

// {
//     "repositories": [
//         "alpine",
//         "webhippie/busybox"
//     ]
// }

// # http GET http://192.168.64.5:5000/v2/alpine/tags/list
// HTTP/1.1 200 OK
// Content-Length: 36
// Content-Type: application/json; charset=utf-8
// Date: Fri, 02 Sep 2016 07:24:18 GMT
// Docker-Distribution-Api-Version: registry/2.0
// X-Content-Type-Options: nosniff

// {
//     "name": "alpine",
//     "tags": [
//         "latest"
//     ]
// }

// # http GET http://192.168.64.5:5000/v2/webhippie/busybox/tags/list
// HTTP/1.1 200 OK
// Content-Length: 47
// Content-Type: application/json; charset=utf-8
// Date: Thu, 01 Sep 2016 22:52:08 GMT
// Docker-Distribution-Api-Version: registry/2.0
// X-Content-Type-Options: nosniff

// {
//     "name": "webhippie/busybox",
//     "tags": [
//         "latest"
//     ]
// }

// # http GET http://192.168.64.5:5000/v2/alpine/manifests/latest
// HTTP/1.1 200 OK
// Content-Length: 5288
// Content-Type: application/vnd.docker.distribution.manifest.v1+prettyjws
// Date: Fri, 02 Sep 2016 07:24:39 GMT
// Docker-Content-Digest: sha256:6a828f13b8595f179cbf7ca249acd5fa98132f9f82e5e62501d6084934361673
// Docker-Distribution-Api-Version: registry/2.0
// Etag: "sha256:6a828f13b8595f179cbf7ca249acd5fa98132f9f82e5e62501d6084934361673"
// X-Content-Type-Options: nosniff

// {
//    "schemaVersion": 1,
//    "name": "alpine",
//    "tag": "latest",
//    "architecture": "amd64",
//    "fsLayers": [
//       {
//          "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
//       },
//       {
//          "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
//       },
//       {
//          "blobSum": "sha256:ff7d71164ebe31fcbf939ce6ba7dc5e0bbb85e400e9f2042d49a72c33d7876ea"
//       },
//       {
//          "blobSum": "sha256:f3c372bbe6931e34e56c42ba318f826993e256facbd329d113f3bec80926ef00"
//       },
//       {
//          "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
//       },
//       {
//          "blobSum": "sha256:e587fa4f6e1fe3d46e8631927252e3f9df509aeb1c14a9cdaabe137e0f78cf24"
//       }
//    ],
//    "history": [
//       {
//          "v1Compatibility": "{\"architecture\":\"amd64\",\"author\":\"Thomas Boerger \\u003cthomas@webhippie.de\\u003e\",\"config\":{\"Hostname\":\"bff8ad45452a\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\",\"TERM=xterm\"],\"Cmd\":[\"bash\"],\"ArgsEscaped\":true,\"Image\":\"sha256:f2d3082562ff69497dd69e099ed3980b0e177a0f08e158ab8cd6293c6db1721a\",\"Volumes\":null,\"WorkingDir\":\"\",\"Entrypoint\":null,\"OnBuild\":[],\"Labels\":{}},\"container\":\"fbc1d39298fc173c8957e2851f68281797195092a1793a6f9c8e131f2596f660\",\"container_config\":{\"Hostname\":\"bff8ad45452a\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\",\"TERM=xterm\"],\"Cmd\":[\"/bin/sh\",\"-c\",\"#(nop) CMD [\\\"bash\\\"]\"],\"ArgsEscaped\":true,\"Image\":\"sha256:f2d3082562ff69497dd69e099ed3980b0e177a0f08e158ab8cd6293c6db1721a\",\"Volumes\":null,\"WorkingDir\":\"\",\"Entrypoint\":null,\"OnBuild\":[],\"Labels\":{}},\"created\":\"2016-06-23T21:26:03.275965535Z\",\"docker_version\":\"1.11.1\",\"id\":\"a905b24213c2313c5704c7170a71d54ba0ef8b6e19e81dbf611b4a43c94d6c90\",\"os\":\"linux\",\"parent\":\"50914605bcb65c1278eab45dd2a573370f159f43548a8f405d286d1fb2d129d7\",\"throwaway\":true}"
//       },
//       {
//          "v1Compatibility": "{\"id\":\"50914605bcb65c1278eab45dd2a573370f159f43548a8f405d286d1fb2d129d7\",\"parent\":\"c5e0e91c6b1d2f5f3ec800925651346349811c2e52cf91b021326ca3427b06aa\",\"created\":\"2016-06-23T21:26:02.298889436Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) ENV TERM=xterm\"]},\"author\":\"Thomas Boerger \\u003cthomas@webhippie.de\\u003e\",\"throwaway\":true}"
//       },
//       {
//          "v1Compatibility": "{\"id\":\"c5e0e91c6b1d2f5f3ec800925651346349811c2e52cf91b021326ca3427b06aa\",\"parent\":\"892e54043a95ada10a36a8bf0d18c79d6ef0777e90681995d44877e58fa7e1a0\",\"created\":\"2016-06-23T21:26:01.363622697Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c apk update \\u0026\\u0026   apk upgrade \\u0026\\u0026   apk add     ca-certificates     curl     bash     bash-completion     ncurses     vim     gettext     logrotate     tar     rsync     shadow@testing     s6@testing \\u0026\\u0026   rm -rf /var/cache/apk/* \\u0026\\u0026   mkdir -p /etc/logrotate.docker.d\"]},\"author\":\"Thomas Boerger \\u003cthomas@webhippie.de\\u003e\"}"
//       },
//       {
//          "v1Compatibility": "{\"id\":\"892e54043a95ada10a36a8bf0d18c79d6ef0777e90681995d44877e58fa7e1a0\",\"parent\":\"563ea9db5fa6423e4cb250489e0f8722a2c07f4c0018f3b945d0b966d8c55b19\",\"created\":\"2016-06-23T21:25:52.641585321Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) ADD dir:93c9e7c188b7c36a9a35eafa6019bc1091fbbc65152719f06dd799182f11b075 in /\"]},\"author\":\"Thomas Boerger \\u003cthomas@webhippie.de\\u003e\"}"
//       },
//       {
//          "v1Compatibility": "{\"id\":\"563ea9db5fa6423e4cb250489e0f8722a2c07f4c0018f3b945d0b966d8c55b19\",\"parent\":\"8e1c48e9ff4b6c2db5a66740f63b7f04e6649b82eae583149d75ef70fb42eb0f\",\"created\":\"2016-06-23T21:25:52.056900266Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) MAINTAINER Thomas Boerger \\u003cthomas@webhippie.de\\u003e\"]},\"author\":\"Thomas Boerger \\u003cthomas@webhippie.de\\u003e\",\"throwaway\":true}"
//       },
//       {
//          "v1Compatibility": "{\"id\":\"8e1c48e9ff4b6c2db5a66740f63b7f04e6649b82eae583149d75ef70fb42eb0f\",\"created\":\"2016-06-23T19:55:23.642524779Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) ADD file:5300ce254c0bf1d3bf6578c41900f3ad965b26a3bf435a3c07860ffc2ae6f7e2 in /\"]}}"
//       }
//    ],
//    "signatures": [
//       {
//          "header": {
//             "jwk": {
//                "crv": "P-256",
//                "kid": "DLUT:TENZ:2RCC:Q3UE:4ZQ3:F7O4:JPYW:XXRK:3ENX:KWZO:HVLW:2APH",
//                "kty": "EC",
//                "x": "poJjoccViNB7hkDx8GpZ1zmjm8uYGaRKsg5MlTfemzo",
//                "y": "QUvGzyLoL_pTETLj3YlumxsHqTggUAo2bU6vBa79698"
//             },
//             "alg": "ES256"
//          },
//          "signature": "OzyQWeq08fqeVfUUte1MS_5U9g3QemW0hoTPD2KYNlotcSb7uUsXXi7eWzohjUh2Wb-XCRpyCM5dtamLcbO_BA",
//          "protected": "eyJmb3JtYXRMZW5ndGgiOjQ2NDEsImZvcm1hdFRhaWwiOiJDbjAiLCJ0aW1lIjoiMjAxNi0wOS0wMlQwNzoyNDozOVoifQ"
//       }
//    ]
// }

// # http GET http://192.168.64.5:5000/v2/webhippie/busybox/manifests/latest
// HTTP/1.1 200 OK
// Content-Length: 2738
// Content-Type: application/vnd.docker.distribution.manifest.v1+prettyjws
// Date: Thu, 01 Sep 2016 22:52:57 GMT
// Docker-Content-Digest: sha256:7783484213b09ebeb87b482bd3c50f85976a882ace875e266997d48c5cd9bc59
// Docker-Distribution-Api-Version: registry/2.0
// Etag: "sha256:7783484213b09ebeb87b482bd3c50f85976a882ace875e266997d48c5cd9bc59"
// X-Content-Type-Options: nosniff

// {
//    "schemaVersion": 1,
//    "name": "webhippie/busybox",
//    "tag": "latest",
//    "architecture": "amd64",
//    "fsLayers": [
//       {
//          "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
//       },
//       {
//          "blobSum": "sha256:8ddc19f16526912237dd8af81971d5e4dd0587907234be2b83e249518d5b673f"
//       }
//    ],
//    "history": [
//       {
//          "v1Compatibility": "{\"architecture\":\"amd64\",\"config\":{\"Hostname\":\"55cd1f8f6e5b\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\"],\"Cmd\":[\"sh\"],\"Image\":\"sha256:e732471cb81a564575aad46b9510161c5945deaf18e9be3db344333d72f0b4b2\",\"Volumes\":null,\"WorkingDir\":\"\",\"Entrypoint\":null,\"OnBuild\":null,\"Labels\":{}},\"container\":\"764ef4448baa9a1ce19e4ae95f8cdd4eda7a1186c512773e56dc634dff208a59\",\"container_config\":{\"Hostname\":\"55cd1f8f6e5b\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\"],\"Cmd\":[\"/bin/sh\",\"-c\",\"#(nop) CMD [\\\"sh\\\"]\"],\"Image\":\"sha256:e732471cb81a564575aad46b9510161c5945deaf18e9be3db344333d72f0b4b2\",\"Volumes\":null,\"WorkingDir\":\"\",\"Entrypoint\":null,\"OnBuild\":null,\"Labels\":{}},\"created\":\"2016-06-23T23:23:37.198943461Z\",\"docker_version\":\"1.10.3\",\"id\":\"b05baf071fd542c3146f08e5f20ad63e76fa4b4bd91f274c4838ddc41f3409f8\",\"os\":\"linux\",\"parent\":\"4185ddbe03f83877b631b5e271a02f6f232de744ae4bfc48ce44216c706cb7fd\",\"throwaway\":true}"
//       },
//       {
//          "v1Compatibility": "{\"id\":\"4185ddbe03f83877b631b5e271a02f6f232de744ae4bfc48ce44216c706cb7fd\",\"created\":\"2016-06-23T23:23:36.73131105Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) ADD file:9ca60502d646bdd815bb51e612c458e2d447b597b95cf435f9673f0966d41c1a in /\"]}}"
//       }
//    ],
//    "signatures": [
//       {
//          "header": {
//             "jwk": {
//                "crv": "P-256",
//                "kid": "DLUT:TENZ:2RCC:Q3UE:4ZQ3:F7O4:JPYW:XXRK:3ENX:KWZO:HVLW:2APH",
//                "kty": "EC",
//                "x": "poJjoccViNB7hkDx8GpZ1zmjm8uYGaRKsg5MlTfemzo",
//                "y": "QUvGzyLoL_pTETLj3YlumxsHqTggUAo2bU6vBa79698"
//             },
//             "alg": "ES256"
//          },
//          "signature": "ZIK936DwN5WoA0r48hra2yRoU8XoN9gbC8zISUeCfZQeSvwECezMBzQ7yjsSg1b1gPKK_SxArZgAwNaMfXFn-g",
//          "protected": "eyJmb3JtYXRMZW5ndGgiOjIwOTEsImZvcm1hdFRhaWwiOiJDbjAiLCJ0aW1lIjoiMjAxNi0wOS0wMVQyMjo1Mjo1N1oifQ"
//       }
//    ]
// }
