## v1.27.1 集群版本命令
```bash
kubectl apply -f - <<EOF
# ------------------- ServiceAccount ------------------- #
apiVersion: v1
kind: ServiceAccount
metadata:
  name: charon
  namespace: kube-system

---
# --------------------- charon ClusterRole --------------------- #
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: charon
rules:
# 核心 API 组（Pods, Services, ConfigMaps, Secrets 等）
- apiGroups: [""]
  resources: ["*"]
  verbs: ["get", "list", "watch", "create", "update", "patch"]
# 扩展 API 组（Deployments, ReplicaSets 等）
- apiGroups: ["apps", "extensions"]
  resources: ["*"]
  verbs: ["get", "list", "watch", "create", "update", "patch"]
# 批量操作 API 组（Jobs, CronJobs）
- apiGroups: ["batch"]
  resources: ["*"]
  verbs: ["get", "list", "watch", "create", "update", "patch"]
# 其他常见 API 组（Ingresses, PersistentVolumes 等）
- apiGroups: ["networking.k8s.io", "storage.k8s.io", "autoscaling"]
  resources: ["*"]
  verbs: ["get", "list", "watch", "create", "update", "patch"]
# 允许访问事件和日志（可选）
- apiGroups: [""]
  resources: ["events", "pods/log"]
  verbs: ["get", "list", "watch"]

---
# ------------- charon Cluster Role Binding ------------- #
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: charon
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: charon
subjects:
- kind: ServiceAccount
  name: charon
  namespace: kube-system
EOF
```