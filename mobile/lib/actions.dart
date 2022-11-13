class Parser {
  Parser(
      {required this.action, required this.reaction, required this.plateforme});
  final List<Reac> reaction;
  final List<Reac> action;
  final String plateforme;
}

class Reac {
  Reac(
      {required this.name,
      required this.description,
      required this.field_names});
  final String name;
  final String description;
  final List<dynamic> field_names;

  factory Reac.fromJson(Map<String, dynamic> data) {
    final name = data['name'] as String;
    final description = data['description'] as String;
    final field_names = data['field_names'] as List<dynamic>;
    return Reac(name: name, description: description, field_names: field_names);
  }
}