namespace RSOI_Lab_01.Entities
{
    public class Person
    {
        /// <summary>
        /// Id юзера
        /// </summary>
        public int? Id { get; set; }
        /// <summary>
        /// Имя юзера
        /// </summary>
        public string? Name { get; set; }
        /// <summary>
        /// Возраст
        /// </summary>
        public int? Age { get; set; }
        /// <summary>
        /// Адрес юзера
        /// </summary>
        public string? Adress { get; set; }
        /// <summary>
        /// Работа юзера
        /// </summary>
        public string? Work { get; set; }

        public Person(int id, string name, int? age = null, string? adress = null, string? work = null) : this(name, age, adress, work)
        {
            Id = id;
        }

        public Person(string name, int? age = null, string? adress = null, string? work = null)
        {
            Name = name;
            Age = age;
            Adress = adress;
            Work = work;
        }

        public override bool Equals(object? obj)
        {
            if (obj as Person is null)
                return false;

            var person = obj as Person;

            return Name == person!.Name && Age == person!.Age && Adress == person!.Adress && Work == person!.Work;
        }

        public override int GetHashCode() => HashCode.Combine(Name, Age, Adress, Work);

        public override string ToString() => $"Person(Id = {Id}, Name = {Name}, Age = {Age}, Adress = {Adress}, Work = {Work})";
    }
}
