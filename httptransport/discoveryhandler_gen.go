// Code generated by openapigen.go DO NOT EDIT.

package httptransport

const (
	_openapiJSON     = `{"components":{"examples":{"Distribution":{"value":{"arch":"","cpe":"","did":"ubuntu","id":"1","name":"Ubuntu","pretty_name":"Ubuntu 18.04.3 LTS","version":"18.04.3 LTS (Bionic Beaver)","version_code_name":"bionic","version_id":"18.04"}},"Environment":{"value":{"distribution_id":"1","introduced_in":"sha256:35c102085707f703de2d9eaad8752d6fe1b8f02b5d2149f1d8357c9cc7fb7d0a","package_db":"var/lib/dpkg/status"}},"Package":{"value":{"arch":"x86","cpe":"","id":"10","kind":"binary","module":"","name":"libapt-pkg5.0","normalized_version":"","source":{"id":"9","kind":"source","name":"apt","source":null,"version":"1.6.11"},"version":"1.6.11"}},"VulnSummary":{"value":{"description":"In the GNU C Library (aka glibc or libc6) before 2.28,\nparse_reg_exp in posix/regcomp.c misparses alternatives,\nwhich allows attackers to cause a denial of service (assertion\nfailure and application exit) or trigger an incorrect result\nby attempting a regular-expression match.\"\n","dist":{"arch":"","cpe":"","did":"ubuntu","id":"0","name":"Ubuntu","pretty_name":"","version":"18.04.3 LTS (Bionic Beaver)","version_code_name":"bionic","version_id":"18.04"},"fixed_in_version":"v0.0.1","links":"http://link-to-advisory","name":"CVE-2009-5155","normalized_severity":"Low","package":{"id":"0","kind":"","name":"glibc","package_db":"","repository_hint":"","source":null,"version":""},"repo":{"id":"0","key":"","name":"Ubuntu 18.04.3 LTS","uri":""}}},"Vulnerability":{"value":{"description":"In the GNU C Library (aka glibc or libc6) before 2.28,\nparse_reg_exp in posix/regcomp.c misparses alternatives,\nwhich allows attackers to cause a denial of service (assertion\nfailure and application exit) or trigger an incorrect result\nby attempting a regular-expression match.\"\n","dist":{"arch":"","cpe":"","did":"ubuntu","id":"0","name":"Ubuntu","pretty_name":"","version":"18.04.3 LTS (Bionic Beaver)","version_code_name":"bionic","version_id":"18.04"},"fixed_in_version":"2.28-0ubuntu1","id":"356835","issued":"2019-10-12T07:20:50.52Z","links":"https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2009-5155\nhttp://people.canonical.com/~ubuntu-security/cve/2009/CVE-2009-5155.html\nhttps://sourceware.org/bugzilla/show_bug.cgi?id=11053\nhttps://debbugs.gnu.org/cgi/bugreport.cgi?bug=22793\nhttps://debbugs.gnu.org/cgi/bugreport.cgi?bug=32806\nhttps://debbugs.gnu.org/cgi/bugreport.cgi?bug=34238\nhttps://sourceware.org/bugzilla/show_bug.cgi?id=18986\"\n","name":"CVE-2009-5155","normalized_severity":"Low","package":{"id":"0","kind":"","name":"glibc","package_db":"","repository_hint":"","source":null,"version":""},"repo":{"id":"0","key":"","name":"Ubuntu 18.04.3 LTS","uri":""},"severity":"Low","updater":""}}},"responses":{"BadRequest":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Error"}}},"description":"Bad Request"},"InternalServerError":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Error"}}},"description":"Internal Server Error"},"MethodNotAllowed":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Error"}}},"description":"Method Not Allowed"},"NotFound":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Error"}}},"description":"Not Found"}},"schemas":{"Callback":{"description":"A callback for clients to retrieve notifications","properties":{"callback":{"description":"the url where notifications can be retrieved","example":"http://clair-notifier/notifier/api/v1/notifications/269886f3-0146-4f08-9bf7-cb1138d48643","type":"string"},"notification_id":{"description":"the unique identifier for this set of notifications","example":"269886f3-0146-4f08-9bf7-cb1138d48643","type":"string"}},"title":"Callback","type":"object"},"Digest":{"description":"A digest string with prefixed algorithm. The format is described here:\nhttps://github.com/opencontainers/image-spec/blob/master/descriptor.md#digests\n\nDigests are used throughout the API to identify Layers and Manifests.\n","example":"sha256:fc84b5febd328eccaa913807716887b3eb5ed08bc22cc6933a9ebf82766725e3","title":"Digest","type":"string"},"Distribution":{"description":"An indexed distribution discovered in a layer. See\nhttps://www.freedesktop.org/software/systemd/man/os-release.html\nfor explanations and example of fields.\n","example":{"$ref":"#/components/examples/Distribution/value"},"properties":{"arch":{"type":"string"},"cpe":{"type":"string"},"did":{"type":"string"},"id":{"description":"A unique ID representing this distribution","type":"string"},"name":{"type":"string"},"pretty_name":{"type":"string"},"version":{"type":"string"},"version_code_name":{"type":"string"},"version_id":{"type":"string"}},"required":["id","did","name","version","version_code_name","version_id","arch","cpe","pretty_name"],"title":"Distribution","type":"object"},"Environment":{"description":"The environment a particular package was discovered in.","properties":{"distribution_id":{"description":"The distribution ID found in an associated IndexReport or\nVulnerabilityReport.\n","example":"1","type":"string"},"introduced_in":{"$ref":"#/components/schemas/Digest"},"package_db":{"description":"The filesystem path or unique identifier of a package database.\n","example":"var/lib/dpkg/status","type":"string"}},"required":["package_db","introduced_in","distribution_id"],"title":"Environment","type":"object"},"Error":{"description":"A general error schema returned when status is not 200 OK","properties":{"code":{"description":"a code for this particular error","type":"string"},"message":{"description":"a message with further detail","type":"string"}},"title":"Error","type":"object"},"IndexReport":{"description":"A report of the Index process for a particular manifest. A\nclient's usage of this is largely information. Clair uses this\nreport for matching Vulnerabilities.\n","properties":{"distributions":{"additionalProperties":{"$ref":"#/components/schemas/Distribution"},"description":"A map of Distribution objects keyed by their Distribution.id\ndiscovered in the manifest.\n","example":{"1":{"$ref":"#/components/examples/Distribution/value"}},"type":"object"},"environments":{"additionalProperties":{"items":{"$ref":"#/components/schemas/Environment"},"type":"array"},"description":"A map of lists containing Environment objects keyed by the\nassociated Package.id.\n","example":{"10":[{"distribution_id":"1","introduced_in":"sha256:35c102085707f703de2d9eaad8752d6fe1b8f02b5d2149f1d8357c9cc7fb7d0a","package_db":"var/lib/dpkg/status"}]},"type":"object"},"err":{"description":"An error message on event of unsuccessful index","example":"","type":"string"},"manifest_hash":{"$ref":"#/components/schemas/Digest"},"packages":{"additionalProperties":{"$ref":"#/components/schemas/Package"},"description":"A map of Package objects indexed by Package.id","example":{"10":{"$ref":"#/components/examples/Package/value"}},"type":"object"},"state":{"description":"The current state of the index operation","example":"IndexFinished","type":"string"},"success":{"description":"A bool indicating succcessful index","example":true,"type":"boolean"}},"required":["manifest_hash","state","packages","distributions","environments","success","err"],"title":"IndexReport","type":"object"},"Layer":{"description":"A Layer within a Manifest and where Clair may retrieve it.","properties":{"hash":{"$ref":"#/components/schemas/Digest"},"headers":{"additionalProperties":{"items":{"type":"string"},"type":"array"},"description":"map of arrays of header values keyed by header\nvalue. e.g. map[string][]string\n","type":"object"},"uri":{"description":"A URI describing where the layer may be found. Implementations\nMUST support http(s) schemes and MAY support additional\nschemes.\n","example":"https://storage.example.com/blob/2f077db56abccc19f16f140f629ae98e904b4b7d563957a7fc319bd11b82ba36\n","type":"string"}},"required":["hash","uri","headers"],"title":"Layer","type":"object"},"Manifest":{"description":"A Manifest representing a container. The 'layers' array must\npreserve the original container's layer order for accurate usage.\n","properties":{"hash":{"$ref":"#/components/schemas/Digest"},"layers":{"items":{"$ref":"#/components/schemas/Layer"},"type":"array"}},"required":["hash","layers"],"title":"Manifest","type":"object"},"Notification":{"description":"A notification expressing a change in a manifest affected by a vulnerability","properties":{"id":{"description":"a unique identifier for this notification","example":"5e4b387e-88d3-4364-86fd-063447a6fad2","type":"string"},"manifest":{"description":"the hash of the manifest affected by the provided vulnerability","example":"sha256:35c102085707f703de2d9eaad8752d6fe1b8f02b5d2149f1d8357c9cc7fb7d0a","type":"string"},"reason":{"description":"the reason for the notifcation, [added | removed]","example":"added","type":"string"},"vulnerability":{"$ref":"#/components/schemas/VulnSummary"}},"title":"Notification","type":"object"},"Package":{"description":"A package discovered by indexing a Manifest","example":{"$ref":"#/components/examples/Package/value"},"properties":{"arch":{"description":"The package's target system architecture","type":"string"},"cpe":{"description":"A CPE identifying the package","type":"string"},"id":{"description":"A unique ID representing this package","type":"string"},"kind":{"description":"Kind of package. Source | Binary","type":"string"},"module":{"description":"A module further defining a namespace for a package","type":"string"},"name":{"description":"Name of the Package","type":"string"},"normalized_version":{"$ref":"#/components/schemas/Version"},"source":{"$ref":"#/components/schemas/SourcePackage"},"version":{"description":"Version of the Package","type":"string"}},"required":["id","name","version"],"title":"Package","type":"object"},"Page":{"description":"A page object indicating to the client how to retrieve multiple pages of a particular entity.","properties":{"next":{"description":"The next id to submit to the api to continue paging","example":"1b4d0db2-e757-4150-bbbb-543658144205","type":"string"},"size":{"description":"The maximum number of elements in a page","example":1,"type":"int"}},"title":"Page"},"PagedNotifications":{"description":"A page object followed by a list of notifications","properties":{"notifications":{"description":"A list of notifications within this page","items":{"$ref":"#/components/schemas/Notification"},"type":"array"},"page":{"description":"A page object informing the client the next page to retrieve.\nIf page.next becomes \"-1\" the client should stop paging.\n","example":{"next":"1b4d0db2-e757-4150-bbbb-543658144205","size":100},"type":"object"}},"title":"PagedNotifications","type":"object"},"Repository":{"description":"A package repository","properties":{"cpe":{"type":"string"},"id":{"type":"string"},"key":{"type":"string"},"name":{"type":"string"},"uri":{"type":"string"}},"title":"Repository","type":"object"},"SourcePackage":{"description":"A source package affiliated with a Package","example":{"$ref":"#/components/examples/Package/value"},"properties":{"arch":{"type":"string"},"cpe":{"description":"A CPE identifying the package","type":"string"},"id":{"description":"A unique ID representing this package","type":"string"},"kind":{"description":"Kind of package. Source | Binary","type":"string"},"module":{"type":"string"},"name":{"description":"Name of the Package","type":"string"},"normalized_version":{"$ref":"#/components/schemas/Version"},"source":{"type":"string"},"version":{"description":"Version of the Package","type":"string"}},"required":["id","name","version"],"title":"SourcePackage","type":"object"},"State":{"description":"an opaque identifier","example":{"state":"aae368a064d7c5a433d0bf2c4f5554cc"},"properties":{"state":{"description":"an opaque identifier","type":"string"}},"required":["state"],"title":"State","type":"object"},"Version":{"description":"Version is a normalized claircore version, composed of a \"kind\" and an\narray of integers such that two versions of the same kind have the\ncorrect ordering when the integers are compared pair-wise.\n","example":"pep440:0.0.0.0.0.0.0.0.0","title":"Version","type":"string"},"VulnSummary":{"description":"A summary of a vulnerability","properties":{"description":{"description":"the vulnerability name","example":"In the GNU C Library (aka glibc or libc6) before 2.28,\nparse_reg_exp in posix/regcomp.c misparses alternatives,\nwhich allows attackers to cause a denial of service (assertion\nfailure and application exit) or trigger an incorrect result\nby attempting a regular-expression match.\"\n","type":"string"},"distribution":{"$ref":"#/components/schemas/Distribution"},"fixed_in_version":{"description":"the version at which the vulnerability is fixed in. empty if not fixed.","example":"v0.0.1","type":"string"},"links":{"description":"links to external information about vulnerability","example":"http://link-to-advisory","type":"string"},"name":{"description":"the vulnerability name","example":"CVE-2009-5155","type":"string"},"normalized_severity":{"description":"A well defined set of severity strings guaranteed to be present.\n","enum":["Unknown","Negligible","Low","Medium","High","Critical","Defcon1"],"type":"string"},"package":{"$ref":"#/components/schemas/Package"},"repository":{"$ref":"#/components/schemas/Repository"}},"title":"VulnSummary","type":"object"},"Vulnerability":{"description":"A unique vulnerability indexed by Clair","example":{"$ref":"#/components/examples/Vulnerability/value"},"properties":{"description":{"description":"A description of this specific vulnerability.","type":"string"},"distribution":{"$ref":"#/components/schemas/Distribution"},"fixed_in_version":{"description":"A unique ID representing this vulnerability.","type":"string"},"id":{"description":"A unique ID representing this vulnerability.","type":"string"},"issued":{"description":"The timestamp in which the vulnerability was issued\n","type":"string"},"links":{"description":"A space separate list of links to any external information.\n","type":"string"},"name":{"description":"Name of this specific vulnerability.","type":"string"},"normalized_severity":{"description":"A well defined set of severity strings guaranteed to be present.\n","enum":["Unknown","Negligible","Low","Medium","High","Critical","Defcon1"],"type":"string"},"package":{"$ref":"#/components/schemas/Package"},"range":{"description":"The range of package versions that are affected by this vulnerability\n","type":"string"},"repository":{"$ref":"#/components/schemas/Repository"},"severity":{"description":"A severity keyword taken verbatim from the vulnerability source.\n","type":"string"},"updater":{"description":"A unique ID representing this vulnerability.","type":"string"}},"required":["id","updater","name","description","links","severity","normalized_severity","fixed_in_version"],"title":"Vulnerability","type":"object"},"VulnerabilityReport":{"description":"A report expressing discovered packages, package environments,\nand package vulnerabilities within a Manifest.\n","properties":{"distributions":{"additionalProperties":{"$ref":"#/components/schemas/Distribution"},"description":"A map of Distribution objects indexed by Distribution.id.\n","example":{"1":{"$ref":"#/components/examples/Distribution/value"}},"type":"object"},"environments":{"additionalProperties":{"items":{"$ref":"#/components/schemas/Environment"},"type":"array"},"description":"A mapping of Environment lists indexed by Package.id","example":{"10":[{"distribution_id":"1","introduced_in":"sha256:35c102085707f703de2d9eaad8752d6fe1b8f02b5d2149f1d8357c9cc7fb7d0a","package_db":"var/lib/dpkg/status"}]},"type":"object"},"manifest_hash":{"$ref":"#/components/schemas/Digest"},"package_vulnerabilities":{"additionalProperties":{"items":{"type":"string"},"type":"array"},"description":"A mapping of Vulnerability.id lists indexed by Package.id.\n","example":{"10":["356835"]}},"packages":{"additionalProperties":{"$ref":"#/components/schemas/Package"},"description":"A map of Package objects indexed by Package.id","example":{"10":{"$ref":"#/components/examples/Package/value"}},"type":"object"},"vulnerabilities":{"additionalProperties":{"$ref":"#/components/schemas/Vulnerability"},"description":"A map of Vulnerabilities indexed by Vulnerability.id","example":{"356835":{"$ref":"#/components/examples/Vulnerability/value"}},"type":"object"}},"required":["manifest_hash","packages","distributions","environments","vulnerabilities","package_vulnerabilities"],"title":"VulnerabilityReport","type":"object"}}},"info":{"contact":{"email":"quay-devel@redhat.com","name":"Clair Team","url":"http://github.com/quay/clair"},"description":"ClairV4 is a set of cooperating microservices which scan, index, and\nmatch your container's content with known vulnerabilities.\n","license":{"name":"Apache License 2.0","url":"http://www.apache.org/licenses/"},"termsOfService":"","title":"ClairV4","version":"0.1"},"openapi":"3.0.2","paths":{"indexer/api/v1/index_report":{"post":{"description":"By submitting a Manifest object to this endpoint Clair will fetch the\nlayers, scan each layer's contents, and provide an index of discovered\npackages, repository and distribution information.\n","operationId":"Index","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Manifest"}}},"required":true},"responses":{"201":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/IndexReport"}}},"description":"IndexReport Created"},"400":{"$ref":"#/components/responses/BadRequest"},"405":{"$ref":"#/components/responses/MethodNotAllowed"},"500":{"$ref":"#/components/responses/InternalServerError"}},"summary":"Index the contents of a Manifest","tags":["Indexer"]}},"indexer/api/v1/index_report/{manifest_hash}":{"get":{"description":"Given a Manifest's content addressable hash an IndexReport will\nbe retrieved if exists.\n","operationId":"GetIndexReport","parameters":[{"description":"A digest of a manifest that has been indexed previous to this\nrequest.\n","in":"path","name":"manifest_hash","required":true,"schema":{"$ref":"#/components/schemas/Digest"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/IndexReport"}}},"description":"IndexReport retrieved"},"400":{"$ref":"#/components/responses/BadRequest"},"404":{"$ref":"#/components/responses/NotFound"},"405":{"$ref":"#/components/responses/MethodNotAllowed"},"500":{"$ref":"#/components/responses/InternalServerError"}},"summary":"Retrieve an IndexReport for the given Manifest hash if exists.","tags":["Indexer"]}},"indexer/api/v1/index_state":{"get":{"description":"The index state endpoint returns a json structure indicating the\nindexer's internal configuration state.\n\nA client may be interested in this as a signal that manifests may need\nto be re-indexed.\n","operationId":"IndexState","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/State"}}},"description":"Indexer State","headers":{"Etag":{"description":"Entity Tag","schema":{"type":"string"}}}},"304":{"description":"Indexer State Unchanged"}},"summary":"Report the indexer's internal configuration and state.","tags":["Indexer"]}},"matcher/api/v1/vulnerability_report/{manifest_hash}":{"get":{"description":"Given a Manifest's content addressable hash a VulnerabilityReport\nwill be created. The Manifest **must** have been Indexed first\nvia the Index endpoint.\n","operationId":"GetVulnerabilityReport","parameters":[{"description":"A digest of a manifest that has been indexed previous to this\nrequest.\n","in":"path","name":"manifest_hash","required":true,"schema":{"$ref":"#/components/schemas/Digest"}}],"responses":{"201":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/VulnerabilityReport"}}},"description":"VulnerabilityReport Created"},"400":{"$ref":"#/components/responses/BadRequest"},"404":{"$ref":"#/components/responses/NotFound"},"405":{"$ref":"#/components/responses/MethodNotAllowed"},"500":{"$ref":"#/components/responses/InternalServerError"}},"summary":"Retrieve a VulnerabilityReport for a given manifest's content\naddressable hash.\n","tags":["Matcher"]}},"notifier/api/v1/notification/{notification_id}":{"delete":{"description":"Issues a delete of the provided notification id and all associated notifications.\nAfter this delete clients will no longer be able to retrieve notifications.\n","operationId":"notifications delete","parameters":[{"description":"A notification ID returned by a callback","in":"path","name":"notification_id","schema":{"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"status":"ok"}}},"description":"OK"},"400":{"$ref":"#/components/responses/BadRequest"},"405":{"$ref":"#/components/responses/MethodNotAllowed"},"500":{"$ref":"#/components/responses/InternalServerError"}},"tags":["Notifier"]},"get":{"description":"By performing a GET with a notification_id as a path parameter the client\nwill retrieve a paginated response of notifcation objects\n","operationId":"notifications get","parameters":[{"description":"A notification ID returned by a callback","in":"path","name":"notification_id","schema":{"type":"string"}},{"description":"The maximum number of notifications to deliver in a single page.","in":"query","name":"page_size","schema":{"type":"int"}},{"description":"The next page to fetch via id. Typically this number is provided to you \non initial response in the page.next field.\nThe first GET request may omit this field.\n","in":"query","name":"next","schema":{"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/PagedNotifications"}}},"description":"A paginated list of notifications"},"400":{"$ref":"#/components/responses/BadRequest"},"405":{"$ref":"#/components/responses/MethodNotAllowed"},"500":{"$ref":"#/components/responses/InternalServerError"}},"summary":"Retreive a paginated result of notifications for the provided id","tags":["Notifier"]}}}}`
	_openapiJSONEtag = `"320da4cfc6a78e15f7b856644aa7a1bca9bc2dc010133db3ea44df500e36af6e"`
)
