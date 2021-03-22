from rest_framework.serializers import ModelSerializer
from .models import Todo, Category


class TodoSerializer(ModelSerializer):

    class Meta:
        model = Todo
        fields = ('id', 'text', 'done')

    
    def update(self, instance: Todo, validated_data):
        instance.text = validated_data.get('text', instance.text)
        instance.done = validated_data.get('done', instance.done)
        instance.save()
        return super().update(instance, validated_data)
    
    def create(self, validated_data):
        todo = Todo.objects.create(
            text=validated_data['text'],
            done=validated_data.get('done', False)
        )

        return todo


class CategorySerializer(ModelSerializer):

    class Meta:
        model = Category
        fields = '__all__'
    
    def create(self, validated_data):
        category = Category.objects.create(
            name=validated_data['name'],
            color=validated_data['color'],
            icon=validated_data.get('icon', None)
        )
        return category