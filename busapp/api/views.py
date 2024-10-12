# api/views.py
from rest_framework import viewsets, permissions
from .models import Ruta, Parada, Buseta, Favorito, Notificacion
from .serializers import (
    RutaSerializer, 
    ParadaSerializer, 
    BusetaSerializer, 
    FavoritoSerializer, 
    NotificacionSerializer, 
    UserSerializer
)
from django.contrib.auth.models import User
from rest_framework.decorators import action
from rest_framework.response import Response
from django.shortcuts import get_object_or_404
from geopy.distance import geodesic

class RutaViewSet(viewsets.ModelViewSet):
    queryset = Ruta.objects.all()
    serializer_class = RutaSerializer
    permission_classes = [permissions.IsAuthenticatedOrReadOnly]

class ParadaViewSet(viewsets.ModelViewSet):
    queryset = Parada.objects.all()
    serializer_class = ParadaSerializer
    permission_classes = [permissions.IsAuthenticatedOrReadOnly]

class BusetaViewSet(viewsets.ModelViewSet):
    queryset = Buseta.objects.all()
    serializer_class = BusetaSerializer
    permission_classes = [permissions.IsAuthenticatedOrReadOnly]

class FavoritoViewSet(viewsets.ModelViewSet):
    queryset = Favorito.objects.all()
    serializer_class = FavoritoSerializer
    permission_classes = [permissions.IsAuthenticated]

    def get_queryset(self):
        # Solo mostrar los favoritos del usuario autenticado
        return Favorito.objects.filter(usuario=self.request.user)

    def perform_create(self, serializer):
        # Asignar el usuario autenticado al crear un favorito
        serializer.save(usuario=self.request.user)

class NotificacionViewSet(viewsets.ModelViewSet):
    queryset = Notificacion.objects.all()
    serializer_class = NotificacionSerializer
    permission_classes = [permissions.IsAuthenticated]

    def get_queryset(self):
        # Solo mostrar las notificaciones del usuario autenticado
        return Notificacion.objects.filter(usuario=self.request.user)

class UserViewSet(viewsets.ReadOnlyModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer
    permission_classes = [permissions.IsAdminUser]

class RutaViewSet(viewsets.ModelViewSet):
    queryset = Ruta.objects.all()
    serializer_class = RutaSerializer
    permission_classes = [permissions.IsAuthenticatedOrReadOnly]

    @action(detail=True, methods=['get'])
    def estimar_tiempo_llegada(self, request, pk=None):
        ruta = self.get_object()
        parada_id = request.query_params.get('parada_id')
        if not parada_id:
            return Response({"error": "parada_id es requerido."}, status=400)
        
        parada = get_object_or_404(Parada, pk=parada_id, ruta=ruta)
        # Obtener la buseta en ruta más cercana a la parada
        busetas = ruta.busetas.filter(estado='en_ruta').order_by('id')  # Puedes mejorar este filtro
        if not busetas.exists():
            return Response({"error": "No hay busetas en ruta."}, status=404)
        
        # Para simplificar, tomamos la primera buseta
        buseta = busetas.first()
        if not buseta.latitud_actual or not buseta.longitud_actual:
            return Response({"error": "Ubicación de la buseta no disponible."}, status=404)
        
        # Calcular la distancia entre la buseta y la parada
        bus_location = (float(buseta.latitud_actual), float(buseta.longitud_actual))
        parada_location = (float(parada.latitud), float(parada.longitud))
        distancia_metros = geodesic(bus_location, parada_location).meters
        
        # Suponiendo una velocidad promedio de 40 km/h (aproximadamente 11.11 m/s)
        velocidad_promedio_m_s = 11.11
        tiempo_segundos = distancia_metros / velocidad_promedio_m_s
        tiempo_minutos = tiempo_segundos / 60
        
        return Response({
            "ruta_id": ruta.id,
            "ruta_nombre": ruta.nombre,
            "parada_id": parada.id,
            "parada_nombre": parada.nombre,
            "distancia_metros": distancia_metros,
            "tiempo_estimado_minutos": round(tiempo_minutos, 2)
        })

