---
components:
  - product-sso
products:
  - name: rhoam
    environments:
      - osd-post-upgrade
    targets:
      - 0.1.0
      - 0.2.0
      - 1.0.0
      - 1.3.0
      - 1.6.0
      - 1.9.0
estimate: 1h
tags:
  - destructive
---

# J08B - Verify User SSO Backup and Restore

## Description

Note: this test should only be performed at a time it will not affect other ongoing testing, or on a separate cluster.
In case the current cluster is not available for some reason, you can perform this test case on the "fresh install" cluster

## Steps

### Postgres

1. Login via `oc` as **kubeadmin**

2. Verify Clients and Realms exist in postgres.

Create a throwaway Postgres instance to access the User SSO Postgres instance

```sh
cat << EOF | oc create -f - -n redhat-rhoam-operator
  apiVersion: integreatly.org/v1alpha1
  kind: Postgres
  metadata:
    name: throw-away-postgres
    labels:
      productName: productName
  spec:
    secretRef:
      name: throw-away-postgres-sec
    tier: development
    type: workshop
EOF
```

```
# Wait until the throwaway Postgres instance is running
oc get pods -n redhat-rhoam-operator | grep throw-away | awk '{print $3}'
# oc rsh to the pod
oc rsh -n redhat-rhoam-operator $(oc get pods -n redhat-rhoam-operator | grep throw-away | awk '{print $1}')
# password and host retrieved from rhssouser-postgres-rhoam secret in redhat-rhoam-operator, psql will prompt for password
psql --host=<db_host> --port=5432 --username=postgres --password --dbname=postgres
select * from client;
select * from realm;
```

Once verified. Delete the throwaway Postgres

```sh
oc delete -n redhat-rhoam-operator postgres/throw-away-postgres
```

3. Run the backup and restore script

```sh
cd test/scripts/backup-restore
NS_PREFIX=redhat-rhoam ./j08-verify-user-sso-backup-and-restore.sh | tee test-output.txt
```

4. Wait for the script to finish without errors
5. Verify in the `test-output.txt` log that the test finished successfully.
