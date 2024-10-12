# api/urls.py
from django.urls import include, path
from rest_framework import routers
from .views import (
    RutaViewSet, 
    ParadaViewSet, 
    BusetaViewSet, 
    FavoritoViewSet, 
    NotificacionViewSet, 
    UserViewSet
)
from rest_framework_simplejwt.views import (
    TokenObtainPairView,
    TokenRefreshView,
)

router = routers.DefaultRouter()
router.register(r'rutas', RutaViewSet)
router.register(r'paradas', ParadaViewSet)
router.register(r'busetas', BusetaViewSet)
router.register(r'favoritos', FavoritoViewSet)
router.register(r'notificaciones', NotificacionViewSet)
router.register(r'usuarios', UserViewSet)

urlpatterns = [
    path('', include(router.urls)),
    path('api-auth/', include('rest_framework.urls', namespace='rest_framework')),
    path('token/', TokenObtainPairView.as_view(), name='token_obtain_pair'),  # Ruta para obtener el token
    path('token/refresh/', TokenRefreshView.as_view(), name='token_refresh'),  # Ruta para refrescar el token
]
