using DS_Lab1.Models;
using System.ComponentModel.DataAnnotations;

namespace DS_Lab1.CommunicationElements
{
    public class PersonResponse
    {
        [Required]
        public int Id { get; set; }
        [Required]
        public string Name { get; set; }
        public int? Age { get; set; }
        public string? Adress { get; set; }
        public string? Work { get; set; }

        public PersonResponse(int id, string name, int? age = null, string? adress = null, string? work = null)
        {
            Id = id;
            Name = name;
            Age = age;
            Adress = adress;
            Work = work;
        }

        public PersonResponse(Person person)
        {
            if (person.Id is null || person.Name is null)
                throw new ArgumentNullException(nameof(person), "Person has no id or name");

            Id = person!.Id.Value;
            Name = person!.Name;
            Age = person.Age;
            Adress = person.Adress;
            Work = person.Work;
        }
    }
}
