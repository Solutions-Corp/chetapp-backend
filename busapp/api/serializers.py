from rest_framework import serializers
from .models import Ruta, Parada, Buseta, Favorito, Notificacion
from django.contrib.auth.models import User

class ParadaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Parada
        fields = '__all__'

class RutaSerializer(serializers.ModelSerializer):
    paradas = ParadaSerializer(many=True, read_only=True)

    class Meta:
        model = Ruta
        fields = '__all__'

class BusetaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Buseta
        fields = '__all__'

class FavoritoSerializer(serializers.ModelSerializer):
    class Meta:
        model = Favorito
        fields = '__all__'

class NotificacionSerializer(serializers.ModelSerializer):
    class Meta:
        model = Notificacion
        fields = '__all__'

class UserSerializer(serializers.ModelSerializer):
    favoritos = FavoritoSerializer(many=True, read_only=True)

    class Meta:
        model = User
        fields = ['id', 'username', 'email', 'favoritos']

class RutaSerializer(serializers.ModelSerializer):
    paradas = ParadaSerializer(many=True, read_only=True)

    class Meta:
        model = Ruta
        fields = '__all__'

